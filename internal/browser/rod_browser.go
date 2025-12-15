//go:build rod
// +build rod

package browser

import "errors"

type RodBrowser struct{}

func NewRodBrowser() Browser {
	return &RodBrowser{}
}

func (r *RodBrowser) Open(url string) error {
	return errors.New("Rod browser disabled in POC build")
}

func (r *RodBrowser) Close() error {
	return nil
}
