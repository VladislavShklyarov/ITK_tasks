package main

import "sync"

type ServerMetric struct {
	Name  string  // Название метрики (например, "memory_usage")
	Value float64 // Значение в байтах
}

func fillMetrics() chan *ServerMetric {
	wg := sync.WaitGroup{}
	metrics := []*ServerMetric{
		// Memory metrics
		{Name: "memory_total", Value: 17179869184}, // 16 GB
		{Name: "memory_used", Value: 12884901888},  // 12 GB
		{Name: "memory_free", Value: 4294967296},   // 4 GB
		{Name: "memory_cache", Value: 2147483648},  // 2 GB

		// CPU metrics
		{Name: "cpu_user", Value: 45.2},
		{Name: "cpu_system", Value: 15.8},
		{Name: "cpu_idle", Value: 39.0},

		// Disk metrics
		{Name: "disk_total", Value: 1073741824000},    // 1 TB
		{Name: "disk_used", Value: 536870912000},      // 500 GB
		{Name: "disk_available", Value: 536870912000}, // 500 GB
	}
	ch := make(chan *ServerMetric)

	wg.Go(func() {
		for _, metric := range metrics {
			ch <- metric
		}
		close(ch)
	})
	return ch
}
