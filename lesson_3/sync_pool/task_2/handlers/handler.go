package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"
	pg "task_2/PostgreSQL"
	sp "task_2/syncPool"
)

// HandleCreateDataWithPool godoc
// @Summary Сохраняет в базу пары "имя":"возраст" через sync.Pool
// @Description Принимает JSON с парами "ключ":"значение" (например, "имя":"возраст"), парсит данные в структуру, выделенную из sync.Pool, и сохраняет их в PostgreSQL.
// @Tags CreateData (With Pool)
// @Accept json
// @Produce json
// @Param request body sp.RequestData true "Payload с данными" example({"name":"Alice","age":"25"})
// @Success 200 {string} string "Успешно сохранено"
// @Failure 400 {string} string "Некорректный JSON"
// @Failure 405 {string} string "Разрешён только метод POST"
// @Router /CreatePool [post]
func HandleCreateDataWithPool(postgres *pg.Postgres, pool *sp.SyncPool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST-method allowed", http.StatusMethodNotAllowed)
			return
		}

		req := pool.Get()
		defer pool.Put(req)

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid JSON "+err.Error(), http.StatusBadRequest)
			return
		}
		conn := postgres.GetConn()
		defer postgres.Release(conn)
		fmt.Printf("Processing request on connection %d...\n", conn.ID)
		for key, value := range req.Data {
			postgres.Create(key, value)
		}

		fmt.Println("objects created:", atomic.LoadInt64(&sp.Created))
	}
}

// HandleCreateDataNoPool godoc
// @Summary Сохраняет в базу пары "имя":"возраст"
// @Description Принимает JSON с парами "ключ":"значение" (например, "имя":"возраст") и сохраняет в PostgreSQL
// @Tags CreateData (No Pool)
// @Accept json
// @Produce json
// @Param request body sp.RequestData true "Payload с данными" example({"name":"Alice","age":"25"})
// @Success 200 {string} string "Успешно сохранено"
// @Failure 400 {string} string "Некорректный JSON"
// @Failure 405 {string} string "Разрешён только метод POST"
// @Router /CreateNoPool [post]
func HandleCreateDataNoPool(postgres *pg.Postgres) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST-method allowed", http.StatusMethodNotAllowed)
			return
		}

		req := sp.NewRequestData()

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "invalid JSON "+err.Error(), http.StatusBadRequest)
			return
		}
		conn := postgres.GetConn()
		defer postgres.Release(conn)
		fmt.Printf("Processing request on connection %d...\n", conn.ID)
		for key, value := range req.Data {
			postgres.Create(key, value)
		}
		fmt.Println("objects created:", atomic.LoadInt64(&sp.Created))
	}
}

// HandleGetData godoc
// @Summary Получает возраст по имени
// @Description Принимает query-параметр `key` (имя пользователя), ищет его в базе и возвращает пару "имя":"возраст".
// @Tags GetData
// @Accept json
// @Produce json
// @Param key query string true "Имя пользователя" example("Alice")
// @Success 200 {object} map[string]string "Найденная пара имя:возраст"
// @Failure 400 {string} string "Отсутствует параметр key"
// @Failure 404 {string} string "Пользователь не найден"
// @Failure 405 {string} string "Разрешён только метод GET"
// @Router /Get [get]
func HandleGetData(postgres *pg.Postgres) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			http.Error(w, "only GET-methods allowed", http.StatusMethodNotAllowed)
			return
		}

		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "missing \"key\" query parameter", http.StatusBadRequest)
			return
		}

		conn := postgres.GetConn()
		defer postgres.Release(conn)

		value, err := postgres.GetData(key)
		fmt.Println(value)

		if err != nil {
			http.Error(w, "not found: "+err.Error(), http.StatusNotFound)
		}

		response := map[string]string{
			key: value,
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}
}

// HandleGetAll godoc
// @Summary Получает все пары "имя":"возраст" из базы
// @Description Возвращает всю сохранённую базу данных в формате JSON — карту, где ключом является имя, а значением возраст.
// @Tags GetData
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "Все пары имя:возраст"
// @Failure 405 {string} string "Разрешён только метод GET"
// @Router /GetAll [get]
func HandleGetAll(postgres *pg.Postgres) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "only GET-methods allowed", http.StatusMethodNotAllowed)
			return
		}
		conn := postgres.GetConn()
		defer postgres.Release(conn)

		response := postgres.GetAll()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}
}
