package main

var c = make(chan int)
var a int
var b int

func f() {
	a = 1
	b = <-c
	println(b)
}
func main() {
	go f()
	c <- 0
	println(a)
}
