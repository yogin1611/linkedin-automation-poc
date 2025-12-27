package messaging

import (
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/storage"
)

type Service struct {
	browser browser.Browser
	store   *storage.StateStore
}

func NewService(b browser.Browser, s *storage.StateStore) *Service {
	return &Service{
		browser: b,
		store:   s,
	}
}

func (s *Service) Send(profileURL string, name string) error {
	// Skip if already messaged
	if s.store.HasMessaged(profileURL) {
		return nil
	}

	message := GenerateMessage(name)

	if err := s.browser.SendMessage(profileURL, message); err != nil {
		return err
	}

	s.store.MarkMessaged(profileURL)
	return nil
}
