package takeaction

import (
	"fmt"
	"reflect"
	"unsafe"
)

func appendOne(s *[]int) {
	*s = append(*s, 1)
}

func SlicesDemo() {
	s1 := []int{0, 0, 0, 0}
	//s1 := make([]int, 4, 8) // capacity is twice the initial size
	s2 := s1

	sh1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	sh2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))

	fmt.Printf("Before appendOne:\ns1: %v sh1: %+v\ns2: %v sh2: %+v\n", s1, sh1, s2, sh2)
	appendOne(&s1)
	fmt.Printf("After appendOne:\ns1: %v sh1: %+v\ns2: %v sh2: %+v\n", s1, sh1, s2, sh2)
	s1[0] = 2
	fmt.Printf("BAfter changing s1:\ns1: %v sh1: %+v\ns2: %v sh2: %+v\n", s1, sh1, s2, sh2)
}

/*
Before appendOne:
s1: [0 0 0 0]
s2: [0 0 0 0]
After appendOne:
s1: [0 0 0 0 1]
s2: [0 0 0 0]
After changing s1:
s1: [2 0 0 0 1]
s2: [0 0 0 0]



Before appendOne:
s1: [0 0 0 0]
s2: [0 0 0 0]
After appendOne:
s1: [0 0 0 0 1]
s2: [0 0 0 0]
After changing s1:
s1: [2 0 0 0 1]
s2: [2 0 0 0]

 */