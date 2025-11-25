package main

import (
	"fmt"
	"sync"
)

const (
	MB = 1024 * 1024
)

func main() {

	metrics := fillMetrics()
	API := make(chan *ServerMetric)

	API = Transformer(metrics)

	for metric := range API {
		fmt.Println(metric)
	}

}

func Transformer(metricsCh chan *ServerMetric) chan *ServerMetric {
	wg := sync.WaitGroup{}
	out := make(chan *ServerMetric)

	wg.Go(func() {
		for metric := range metricsCh {
			switch metric.Name {
			case "memory_total", "memory_used", "memory_free", "memory_cache",
				"disk_total", "disk_used", "disk_available",
				"network_rx", "network_tx", "disk_io_read", "disk_io_write":
				metric.Value = metric.Value / 1024 / 1024 // в МБ
				// cpu_user, cpu_system, cpu_idle, error_rate оставляем как есть
			}
			out <- metric
		}
		close(out)
	})
	return out
}
