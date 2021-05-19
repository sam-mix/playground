package crashy

import "github.com/go-errors/errors"

var Crashed = errors.Errorf("oh dear")

func Crash() error {
	return errors.New(Crashed)
}
