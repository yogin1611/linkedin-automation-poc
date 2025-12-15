package main

import (
	"time"

	"linkedin-automation/internal/behavior"
	"linkedin-automation/internal/browser"
	"linkedin-automation/internal/config"
	"linkedin-automation/internal/logger"
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

	// Load persistent state
	store := storage.NewJSONStore("state.json")
	state, err := store.Load()
	if err != nil {
		log.Error("Failed to load state: " + err.Error())
		return
	}

	// Initialize behavior engine
	beh := behavior.NewHumanBehavior()

	// Apply fingerprint strategy (mocked)
	fp := &behavior.MockFingerprint{}
	_ = fp.Apply()

	// Initialize mock browser
	br := browser.NewMockBrowser()

	// ---- Simulated human-like workflow ----

	// Think before action
	beh.Think()

	// Simulate mouse movement
	behavior.SimulateMouseMovement(
		behavior.MousePoint{X: 0, Y: 0},
		behavior.MousePoint{X: 120, Y: 260},
	)

	// Open page (mock)
	if err := br.Open("https://example.com"); err != nil {
		log.Error("Failed to open page: " + err.Error())
		return
	}

	// Simulate scrolling
	behavior.SimulateScrolling()

	// Simulate typing (e.g., connection note / message)
	behavior.SimulateTyping("Hello, thanks for connecting!")

	// Short & long pauses to mimic human pacing
	beh.ShortPause()
	beh.LongPause()

	// Update state
	state.OpenedURLs = append(state.OpenedURLs, "https://example.com")
	state.LastRun = time.Now()

	// Close browser (mock)
	_ = br.Close()

	// Save state
	if err := store.Save(state); err != nil {
		log.Error("Failed to save state: " + err.Error())
		return
	}

	log.Info("State saved. Mock browser closed successfully")
}
