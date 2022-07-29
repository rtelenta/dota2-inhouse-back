package stats

import (
	"os"
	"time"
)

var startTime time.Time

func uptime() string {
	return time.Since(startTime).String()
}

func init() {
	startTime = time.Now()
}

func GetOsStats() map[string]any {
	hostname, _ := os.Hostname()

	return map[string]any{
		"app":      "dota2",
		"hostname": hostname,
		"uptime":   uptime(),
	}
}
