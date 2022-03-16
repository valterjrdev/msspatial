package tests

import "errors"

type Reader int

func (Reader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}
