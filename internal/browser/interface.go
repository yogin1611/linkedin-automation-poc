package browser

type Browser interface {
	Open(url string) error
	Close() error
	SendMessage(profileURL string, message string) error
}
