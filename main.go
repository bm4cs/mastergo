package main

import (
	"github.com/bm4cs/mastergo/web"
	"unicode"
)

func acronym(s string) (acr string) {
	afterSpace := false

	for i, e := range s {
		if (afterSpace || i == 0) && unicode.IsLetter(e) && unicode.IsUpper(e) {
			acr += string(e)
			afterSpace = false
		}
		if unicode.IsSpace(e) {
			afterSpace = true
		}
	}

	return acr
}

func main() {

	//fmt.Println(takeaction.LongestString("Six", "sleek", "swans", "swam", "swiftly", "southwards"))

	//bank.Run()

	web.RunWebServer()

}
