package browser

import (
	"fmt"
	"time"
)

type MockBrowser struct{}

func NewMockBrowser() Browser {
	return &MockBrowser{}
}

func (m *MockBrowser) Open(url string) error {
	fmt.Println("[MOCK BROWSER] Opening URL:", url)
	time.Sleep(2 * time.Second)

	fmt.Println("[MOCK BROWSER] Simulating page load...")
	time.Sleep(2 * time.Second)

	fmt.Println("[MOCK BROWSER] Page loaded successfully")
	return nil
}

func (m *MockBrowser) Close() error {
	fmt.Println("[MOCK BROWSER] Closing browser")
	time.Sleep(1 * time.Second)
	return nil
}

func (m *MockBrowser) SendMessage(profileURL string, message string) error {
	fmt.Println("[MOCK MESSAGE] Sending message")
	fmt.Println("To:", profileURL)
	fmt.Println("Message:", message)
	time.Sleep(1 * time.Second)
	return nil
}
