package storage

import (
	"encoding/json"
	"os"
	"time"
)

type StateStore struct {
	filePath string
	state    *State
}

func NewStateStore(path string) (*StateStore, error) {
	store := &StateStore{
		filePath: path,
		state:    NewState(),
	}

	// Load existing state if present
	if err := store.load(); err != nil {
		return nil, err
	}

	return store, nil
}

func (s *StateStore) load() error {
	file, err := os.Open(s.filePath)
	if err != nil {
		// File doesn’t exist → fresh state
		return nil
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&s.state)
}

func (s *StateStore) Save() error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(s.state)
}

/* =========================
   Messaging helpers
   ========================= */

func (s *StateStore) HasMessaged(profileURL string) bool {
	_, exists := s.state.MessagesSent[profileURL]
	return exists
}

func (s *StateStore) MarkMessaged(profileURL string) {
	s.state.MessagesSent[profileURL] = time.Now().Format(time.RFC3339)
	_ = s.Save()
}
