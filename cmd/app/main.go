package main

import (
	"linkedin-automation/internal/behavior"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/config"
	"linkedin-automation/internal/logger"
	"linkedin-automation/internal/messaging"
	"linkedin-automation/internal/storage"
)

func main() {
	// Load config & logger
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)

	log.Info("Starting LinkedIn Automation POC (Mock Mode)")

	// Activity scheduling (business hours)
	if !behavior.IsWithinBusinessHours() {
		log.Info("Outside business hours. Exiting.")
		return
	}

	// Initialize persistent state store
	store, err := storage.NewStateStore("state.json")
	if err != nil {
		log.Error("Failed to initialize state store: " + err.Error())
		return
	}

	// Initialize behavior engine
	beh := behavior.NewHumanBehavior()

	// Apply fingerprint strategy (mocked)
	fp := &behavior.MockFingerprint{}
	_ = fp.Apply()

	// Initialize mock browser
	br := browser.NewMockBrowser()

	// Initialize messaging service
	msgService := messaging.NewService(br, store)

	// ---- Simulated human-like workflow ----

	// Think before action
	beh.Think()

	// Simulate mouse movement
	behavior.SimulateMouseMovement(
		behavior.MousePoint{X: 0, Y: 0},
		behavior.MousePoint{X: 120, Y: 260},
	)

	// Open page (mock)
	profileURL := "https://linkedin.com/in/sample-profile"
	profileName := "John Doe"

	if err := br.Open(profileURL); err != nil {
		log.Error("Failed to open page: " + err.Error())
		return
	}

	// Simulate scrolling
	behavior.SimulateScrolling()

	// Simulate typing (pre-message interaction)
	behavior.SimulateTyping("Reviewing profile...")

	// Short pause
	beh.ShortPause()

	// ---- Messaging workflow ----
	if err := msgService.Send(profileURL, profileName); err != nil {
		log.Error("Failed to send message: " + err.Error())
		return
	}

	// Long pause after messaging
	beh.LongPause()

	// Close browser (mock)
	_ = br.Close()

	log.Info("Workflow completed successfully (Mock Mode)")
}
