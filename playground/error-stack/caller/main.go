package main

import (
	"fmt"
	"playground/playground/error-stack/crashy"

	"github.com/go-errors/errors"
)

func main() {
	err := crashy.Crash()
	if err != nil {
		if errors.Is(err, crashy.Crashed) {
			fmt.Println(err.(*errors.Error).ErrorStack())
		} else {
			panic(err)
		}
	}
}
