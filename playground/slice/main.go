package main

import "fmt"

func main() {
	s := []uint64{1, 2, 3, 4, 5}

	s1 := s[0:]
	s1[0] = 11

	s2 := s[0:]
	s2 = append(s2, 6)
	s2[0] = 111

	s3 := s[0:1:2]
	fmt.Println(s3)
	s3 = append(s3, 7)

	fmt.Println(s)
}
