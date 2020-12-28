package takeaction

import "fmt"

func PointersNew() {
	p := new(int32)
	*p = 64
	fmt.Println("p's value is a's address", p)
	fmt.Println("p's value yields a's value", *p)
}

func PointersBasics() {
	var a int = 1337
	var p *int
	p = &a
	fmt.Println("p's value is a's address", p)
	fmt.Println("p's value yields a's value", *p)
}