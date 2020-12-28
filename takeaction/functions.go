package takeaction

import (
	"fmt"
	"unicode"
)

func LongestString(s ...string) int {
	longest := 0
	var p *string

	for _, s := range s {
		if len(s) > longest {
			longest = len(s)
			p = &s
		}
	}

	fmt.Printf("The longest string is %s\n", *p)
	return longest
}

func ScopeMadness() {
	str := "abcde"
	for _, ch := range str {
		chu := unicode.ToUpper(ch)
		fmt.Print(string(chu))
	}
	fmt.Println("\n" + str)
}

func newClosure() func() int {
	var a int
	return func() int {
		a++
		fmt.Println(a)
		return a
	}
}

// return two closures
// is the outer scope state shared between the closures??
func newTwoClosures() (func(), func() int) {
	var foo int
	f1 := func() {
		foo = 5
	}
	f2 := func() int {
		foo *= 7
		return foo
	}
	return f1, f2
}

func SingleClosure() int {
	c := newClosure()
	c()
	c()
	return c()
}

func DoubleClosures() int {
	f1, f2 := newTwoClosures()
	f1()
	n := f2()
	return n
}

// Task 2: Clever tracing with “defer”
// https://appliedgo.com/courses/128278/lectures/2701459

func trace(name string) func() {
	fmt.Printf("Entering %s\n", name)
	return func() {
		fmt.Printf("Leaving %s\n", name)
	}
	// TODO:
	// 1. Print "Entering <name>"
	// 2. return a func() that prints "Leaving <name>"
}

func f() {
	defer trace("f")()
	fmt.Println("Doing something")
}

func CleverTracingWithDefer() {
	fmt.Println("Before f")
	f()
	fmt.Println("After f")
}
