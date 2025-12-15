package storage

import "time"

type State struct {
	OpenedURLs []string  `json:"opened_urls"`
	LastRun    time.Time `json:"last_run"`
}
