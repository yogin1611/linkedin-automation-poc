package behavior

import "fmt"

type FingerprintStrategy interface {
	Apply() error
}

type MockFingerprint struct{}

func (m *MockFingerprint) Apply() error {
	fmt.Println("[FINGERPRINT] Applying randomized browser fingerprint (mock)")
	return nil
}
