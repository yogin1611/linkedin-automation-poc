package behavior

import "time"

func IsWithinBusinessHours() bool {
	hour := time.Now().Hour()
	return hour >= 9 && hour <= 20
}
