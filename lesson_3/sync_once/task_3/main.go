package main

import (
	"fmt"
	"log"
	"sync"
	pm "task_3/pluginManager"
)

func initDemo() (pm.Plugin, error) {
	// Имитация длительной инициализации
	// time.Sleep(500 * time.Millisecond)
	return &pm.DemoPlugin{}, nil
}

func main() {
	plugMan := pm.NewPluginManager()

	plugMan.RegisterPlugin("demo", initDemo)
	plugMan.RegisterPlugin("broken", func() (pm.Plugin, error) {
		return nil, fmt.Errorf("simulated error")
	})

	var wg sync.WaitGroup

	// Тестирование рабочего плагина
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p, err := plugMan.GetPlugin("demo")
			if err != nil {
				log.Printf("Goroutine %d error: %v", id, err)
				return
			}
			log.Printf("Goroutine %d: %s", id, p.Execute())
		}(i)
	}

	// Тестирование плагина с ошибкой
	for i := 5; i < 7; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			_, err := plugMan.GetPlugin("broken")
			if err != nil {
				log.Printf("Goroutine %d error: %v", id, err)
			}
		}(i)
	}

	wg.Wait()
}
