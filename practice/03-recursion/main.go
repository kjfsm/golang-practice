package main

import "fmt"

func main() {
	values := map[int]int {
		8: 40320,
		1: 1,
		4: 24,
		6: 720,
		0: 1,
		12: 479001600,
	}

	for key, value := range values {
		answer := factorial(key)
		fmt.Println(answer)
		if answer != value {
			fmt.Println("error")
			return
		}
	}
	fmt.Println("clear!!")
}

func factorial(a int) int {
	// ここを変える
	return a
}
