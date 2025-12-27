package messaging

import "fmt"

func GenerateMessage(name string) string {
	return fmt.Sprintf(
		"Hi %s, I came across your profile and would love to connect and exchange insights.",
		name,
	)
}
