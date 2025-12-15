package browser

type Browser interface {
	Open(url string) error
	Close() error
}
