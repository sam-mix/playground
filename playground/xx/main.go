package main

import "fmt"

type x struct {
	X int
}

func main() {
	xs := []*x{}
	xs = append(xs, &x{})
	xs = append(xs, &x{})
	xs = append(xs, &x{})
	for _, i := range xs {
		i.X = 100
	}
	for _, i := range xs {
		fmt.Println(i.X)
	}
}
