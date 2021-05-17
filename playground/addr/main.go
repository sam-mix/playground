package main

import "fmt"

func main() {
	//s := [3]int{9, 8, 7}
	s := []int{9, 8, 7}
	//p := &s
	//r := *p
	r := s
	r[0] = 11
	fmt.Println(s[0])
}
