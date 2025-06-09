package notifier

import (
	"log"
	"sys-monitor-go/monitor"

	"github.com/go-resty/resty/v2"
)

func SendAlert(api string, stats monitor.Stats) {
	client := resty.New()

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(stats).
		Post(api)

	if err != nil {
		log.Println("Failed to send alert:", err)
	} else {
		log.Println("Alert sent:", resp.Status())
	}
}
