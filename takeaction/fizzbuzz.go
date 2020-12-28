package takeaction

import (
	"fmt"
	// "os"
	"strconv"
)

func fizzbuzz(n int) {

	for i := 1; i < n; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Print(strconv.Itoa(i), " ")
		}
	}
}

// func main2() {
// 	n := 50
// 	if len(os.Args) > 1 {
// 		max, err := strconv.Atoi(os.Args[1])
// 		if err == nil {
// 			n = max
// 		}
// 	}
// 	fizzbuzz(n)
// }
