package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := int64(1)
	x(i)
}

func x(i interface{}) {
	v := reflect.ValueOf(i)
	t := reflect.TypeOf(i)

	fmt.Println(v)
	fmt.Println(t)
}
