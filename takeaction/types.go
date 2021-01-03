package takeaction

import (
	"bufio"
	"fmt"
	"os"
)

// Task 1
// Go doesnt provide a set type
// lets make our own
type set map[string]struct{}

// Task 2
// Your task is to write a function iter(n int) <return value> that takes a number n and returns someting that can be passed to a range loop, like so:
//
//for i := range iter(7) {
//    ...
//}
func StructBasicsIntegerIterator() {
	for i := range iter(7) {
		fmt.Println(i)
	}
}

// ben made this
// returns a slice of empty (0 byte) structs
func iter(i int) []struct{} {
	return make([]struct{}, i)
}

// Task 3: Bonus task: Duplicate finder
//
// Write a small program that
//  opens a file,
//  reads the file line by line, and
//  verifies if the line already exists in the set
//
// If so, it shall print the line; otherwise, it shall add the line to the set.
func NamedTypeSetDemo() {

	words := make(set, 10)
	//words["foo"] = struct{}{}

	f, e := os.Open("assets/poem.txt")
	if e != nil {
		fmt.Printf("could not open file: %s\n", e)
		return
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		word := s.Text()
		_, exists := words[word]

		if exists {
			fmt.Println(word)
		} else {
			words[word] = struct{}{}
		}
	}

}
