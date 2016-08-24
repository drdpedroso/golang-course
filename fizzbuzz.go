package main

import (
				"fmt"
        "strconv"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(x int) string {
	if x % 3 == 0 && x % 5 != 0 {
		return "fizz"
	} else if x % 5 == 0 && x % 3 != 0 {
		return "buzz"
	} else if x % 5 == 0 && x % 3 == 0 {
		return "FizzBuzz"
	}

	t := strconv.Itoa(x)
	return t
}
