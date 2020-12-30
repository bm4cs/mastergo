package takeaction

import (
	"fmt"
	"strings"
)

func countWords(s string, m map[string]int) {
	words := strings.Split(s, " ")

	for _, e := range words {
		var word = strings.ReplaceAll(e, "\n", "")
		word = strings.Trim(word, ".,:;?!()[]{}-—/\\\t\n\"'")
		word = strings.ToLower(word)
		//fmt.Printf("adding word %s\n", word)
		m[word]++
	}
}

func printCount(m map[string]int) {
	for k, v := range m {
		if v > 1 {
			fmt.Printf("%s := %d\n", k, v)
		}
	}
}

func readWords() string {
	return "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand. "
}

func MapDemo() {
	var m map[string]int = make(map[string]int, 50)
	words := readWords()
	countWords(words, m)
	printCount(m)
}
