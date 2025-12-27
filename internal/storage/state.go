package storage

import "time"

type State struct {
	VisitedProfiles []string          `json:"visited_profiles"`
	MessagesSent    map[string]string `json:"messages_sent"`
	LastExecution   time.Time         `json:"last_execution"`
}

func NewState() *State {
	return &State{
		VisitedProfiles: []string{},
		MessagesSent:    map[string]string{},
		LastExecution:   time.Now(),
	}
}
