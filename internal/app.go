package internal

import "github.com/praveenmahasena/aiclient/internal/dialer"

func Run() error {
	d := dialer.New(":42069")
	return d.Start()
}
