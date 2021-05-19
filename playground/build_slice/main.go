package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	length := 1
	var ptr unsafe.Pointer
	// var s1 = struct {
	// 	addr uintptr
	// 	len  int
	// 	cap  int
	// }{
	// 	addr: ptr,
	// 	len:  length,
	// 	cap:  length,
	// }
	// s := *(*[]byte)(unsafe.Pointer(&s1))
	// fmt.Println(s)
	var o []byte
	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&o)))
	sliceHeader.Cap = length
	sliceHeader.Len = length
	sliceHeader.Data = uintptr(ptr)
	fmt.Println(sliceHeader)
}
