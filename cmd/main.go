package main

import (
	"log"
	"net/http"
	"time"

	"sys-monitor-go/monitor"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_cpu_usage_percent",
		Help: "CPU usage percentage",
	})
	memGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_memory_usage_percent",
		Help: "Memory usage percentage",
	})
	diskUsedGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_disk_used_bytes",
		Help: "Disk used in bytes",
	})
	diskTotalGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_disk_total_bytes",
		Help: "Total disk size in bytes",
	})
	diskUsedPercentGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "system_disk_usage_percent",
		Help: "Disk usage percentage",
	})
)

func init() {
	// Đăng ký các metric với Prometheus
	prometheus.MustRegister(cpuGauge, memGauge, diskUsedGauge, diskTotalGauge, diskUsedPercentGauge)
}

func main() {
	// Start HTTP server expose /metrics cho Prometheus scrape
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Println("Starting Prometheus metrics server at :2112/metrics")
		if err := http.ListenAndServe(":2112", nil); err != nil {
			log.Fatal(err)
		}
	}()

	// Thay đổi ticker thành 1 giây
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		stats := monitor.GetStats()

		cpuGauge.Set(stats.CPU * 100) // nhân 100 để ra %
		memGauge.Set(stats.Memory)    // nhân 100 để ra %
		diskUsedGauge.Set(float64(stats.DiskUsed))
		diskTotalGauge.Set(float64(stats.DiskTotal))
		diskUsedPercentGauge.Set(stats.DiskUsedPercent * 100)

		log.Printf("CPU: %.2f%% | RAM: %.2f%% | Disk: %s / %s (%.2f%%)\n",
			stats.CPU*100,
			stats.Memory*100,
			monitor.FormatBytes(stats.DiskUsed),
			monitor.FormatBytes(stats.DiskTotal),
			stats.DiskUsedPercent,
		)
	}
}
