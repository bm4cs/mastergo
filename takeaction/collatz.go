package takeaction

import (
	"fmt"
	"os"
	"strconv"
)

func collatz(n int) int {

	var i int

	for i = 0; n != 1; i++ {

		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
	}

	return i
}

func Run(n int) int {
	return collatz(n)
}

func RunCli() {
	var n int
	var err error

	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1]) //.ParseInt(os.Args[1], 10, 0)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		fmt.Println("Input a number kind sir > 1: ")
		_, err := fmt.Scanln("%d", &n)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if n <= 1 {
		fmt.Println("n must be larger than 1 dude")
		return
	}

	c := collatz(n)
	fmt.Println(n, "requires", c, "steps to reach 1")
}
