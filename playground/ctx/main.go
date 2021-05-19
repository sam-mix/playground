package main

import (
	"context"
	"fmt"
	"time"
)

func r1() {
	c1 := context.Background()
	context.WithValue(c1, 1, 1)
	fmt.Println(c1.Value(1))
}

func r2() {
	c1 := context.WithValue(context.Background(), 1, 1)
	c2, _ := context.WithCancel(c1)
	fmt.Println(c1.Value(1))
	fmt.Println(c2.Value(1))
}

func r3() {
	c1, _ := context.WithTimeout(context.TODO(), time.Second)
	c2, _ := context.WithTimeout(c1, time.Second*2)
	time.Sleep(time.Millisecond * 1100)
	select {
	case <-c2.Done():
		fmt.Println(1)
	default:
		fmt.Println(2)

	}
}

func main() {
	r1()
	r2()
	r3()
}
