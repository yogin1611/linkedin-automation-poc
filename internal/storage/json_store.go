package storage

import (
	"encoding/json"
	"os"
)

type JSONStore struct {
	filePath string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{filePath: path}
}

func (s *JSONStore) Load() (*State, error) {
	file, err := os.Open(s.filePath)
	if err != nil {
		// File doesn’t exist yet → start fresh
		return &State{}, nil
	}
	defer file.Close()

	var state State
	if err := json.NewDecoder(file).Decode(&state); err != nil {
		return nil, err
	}

	return &state, nil
}

func (s *JSONStore) Save(state *State) error {
	file, err := os.Create(s.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(state)
}
