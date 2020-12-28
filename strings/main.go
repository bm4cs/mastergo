package mastergo

import (
	"fmt"
	"os"
	"strings"
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

	fmt.Println("IndexRune:", strings.IndexRune("abä走.", '走'))

	s := "Pan Galactic Gargle Blaster"
	if len(os.Args) > 1 {
		s = strings.Join(os.Args, " ")
	}
	fmt.Println(acronym(s))
}
