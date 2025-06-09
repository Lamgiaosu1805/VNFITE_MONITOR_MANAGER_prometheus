package monitor

import (
	"fmt"
	"log"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type Stats struct {
	CPU             float64 // % CPU sử dụng
	Memory          float64 // % RAM sử dụng
	DiskUsed        uint64  // dung lượng ổ đã dùng (bytes)
	DiskTotal       uint64  // tổng dung lượng ổ (bytes)
	DiskUsedPercent float64 // % dung lượng ổ sử dụng
}

func GetStats() Stats {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		log.Println("Error getting CPU percent:", err)
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error getting memory stats:", err)
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		log.Println("Error getting disk stats:", err)
	}

	return Stats{
		CPU:             cpuPercent[0],
		Memory:          vmStat.UsedPercent,
		DiskUsed:        diskStat.Used,
		DiskTotal:       diskStat.Total,
		DiskUsedPercent: diskStat.UsedPercent,
	}
}

// Hàm helper để format dung lượng byte cho dễ đọc
func FormatBytes(bytes uint64) string {
	const (
		KB = 1 << 10
		MB = 1 << 20
		GB = 1 << 30
		TB = 1 << 40
	)
	switch {
	case bytes >= TB:
		return fmt.Sprintf("%.2f TB", float64(bytes)/TB)
	case bytes >= GB:
		return fmt.Sprintf("%.2f GB", float64(bytes)/GB)
	case bytes >= MB:
		return fmt.Sprintf("%.2f MB", float64(bytes)/MB)
	case bytes >= KB:
		return fmt.Sprintf("%.2f KB", float64(bytes)/KB)
	default:
		return fmt.Sprintf("%d B", bytes)
	}
}
