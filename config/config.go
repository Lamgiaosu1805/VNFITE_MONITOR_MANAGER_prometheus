package config

import "time"

type Config struct {
	CPUThreshold    float64
	MemoryThreshold float64
	DiskThreshold   float64
	AlertAPI        string
	CheckInterval   time.Duration
}

func Load() Config {
	return Config{
		CPUThreshold:    80,
		MemoryThreshold: 80,
		DiskThreshold:   85,
		AlertAPI:        "http://your-api.com/alert",
		CheckInterval:   30 * time.Second,
	}
}
