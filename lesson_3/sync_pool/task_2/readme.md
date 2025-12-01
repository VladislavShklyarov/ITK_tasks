В этом задании попробовал реализовать более-менее полноценный http-сервер,
используя знания из предыдущих заданий.

# API:

`/CreateNoPoll` - принимает запрос JSON, парсит его без sync.Pool и затем отправляет в базу данных.<br>
`/CreatePoll` - то же самое, только с sync.Pool.<br>
`/Get` - в query-params передаем имя и возвращаем значение из базы данных.<br>
`/GetAll` - возвращает все записи, которые есть в базе.<br>

Более подробная документация есть в swagger по адресу `http://localhost:8080/swagger/index.html#/`

# Особенности реализации

1. Вся работа с базами данных осуществляется через очередь подключений,
как в одном из предыдущих задач.<br> В каждом хэндлере мы получаем подключение,
делаем работу и возвращаем его обратно. В main файле количество подключений = 5

В логах командной строки хорошо видно:
```bash
Server started at :8080
Connection 0 acquired
Processing request on connection 0...
objects created: 2
Connection 0 released
req ptr: 0xc000a76180

Connection 1 acquired
Processing request on connection 1...
objects created: 2
Connection 1 released
```
Так работает, если запросы последовательные. Если же мы отправим много
запросов одновременно, то вся очередь заполнится, и горутинам придется
подождать пока не освободится хотя бы одна.

Для отправки таких запросов быстренько сгенерировал клиента `task_2_client`.

```bash
Server started at :8080
Connection 0 acquired
Connection 1 acquired
Connection 2 acquired
Connection 3 acquired
Connection 4 acquired
All connections are busy, waiting...
All connections are busy, waiting...
All connections are busy, waiting...
All connections are busy, waiting...
Processing request on connection 4...
objects created: 18
Connection 4 released
Processing request on connection 0...
Processing request on connection 3...
objects created: 18
Processing request on connection 1...
objects created: 18
Processing request on connection 2...
```
Первые 5 запросов заняли всю очередь, остальные остались ждать.
Затем, по мере освобождения подключений начали обрабатываться.

# 2. Sync.Pool

В sync.Pool я помещаю сами структуры `RequestData`, в который парсится приишедший Json.
Также внутри есть `atomic.AddInt64(&Created, 1)`. Это счетчик, который показывает,
сколько таких объектов лежит в памяти (по крайней мере так задумано).

Тестировал при помощи Postman и отправил 23 последоватлеьных запроса на `/CreatePool`с задержкой 0ms.

```bash
Server started at :8080
Connection 0 acquired
Processing request on connection 0...
objects created: 2
Connection 0 released
req ptr: 0xc0001a43f0
Connection 1 acquired
Processing request on connection 1...
objects created: 2
Connection 1 released
req ptr: 0xc0001a43f0
Connection 2 acquired
Processing request on connection 2...
objects created: 4
Connection 2 released
req ptr: 0xc00008a0c0
Connection 3 acquired
Processing request on connection 3...
objects created: 4
Connection 3 released
req ptr: 0xc00008a0c0
<----....----->
```
По началу создалось 2 объекта, затем выросло до 4-х и больше не добавлялось.
При этом адреса выводятся явно одни и те же, так что я думаю, что оно
работает плюс минус как задумано.

Если отправлять на `/CreateNoPool`, то на каждый будет создан новый объект в памяти:
```bash
Server started at :8080
Connection 0 acquired
Processing request on connection 0...
objects created: 1
Connection 0 releasedы
Connection 1 acquired
Processing request on connection 1...
objects created: 2
Connection 1 released
Connection 2 acquired
Processing request on connection 2...
objects created: 3
Connection 2 released
Connection 3 acquired
Processing request on connection 3...
objects created: 4
Connection 3 released
Connection 4 acquired
Processing request on connection 4...
objects created: 5
Connection 4 released
```

На 5 запросов 5 объектов.

Правда, если отправить даже на `/CreateWithPool` при помощи клиента много запросов,
то в памяти появляется объект на каждый из них:

```bash
req ptr: 0xc000b8a8d0
Connection 0 acquired
Processing request on connection 2...
objects created: 87
Connection 2 released
req ptr: 0xc000b8a9f0
Processing request on connection 1...
objects created: 87
Connection 1 released
req ptr: 0xc000a6c690
Processing request on connection 4...
objects created: 87
Processing request on connection 3...
objects created: 87
Connection 4 released
req ptr: 0xc000b8ab10
Connection 3 released
req ptr: 0xc000a83410
Processing request on connection 0...
objects created: 87
Connection 0 released
req ptr: 0xc000b8ac30
```
Думаю, так происходит из-за того, что запросы приходят в один момент времени,
и программе приходится выделить память под каждый из них.
Возможно чтобы исправить это, можно также придумать очередь, sync.Cond или 
что-нибудь в этом духе.