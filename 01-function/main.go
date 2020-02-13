package main

import (
	"fmt"
)

func main() {
	answer1 := add(3, 4)
	fmt.Println("3 + 4 = ", answer1)
	if answer1 != 7 {
		fmt.Println("error")
		return
	}

	answer2 := sub(9, 4)
	fmt.Println("9 - 4 = ", answer2)
	if answer2 != 5 {
		fmt.Println("error")
		return
	}

	answer3 := mul(3, 4)
	fmt.Println("3 * 4 = ", answer3)
	if answer3 != 12 {
		fmt.Println("error")
		return
	}

	answer4 := div(15, 3)
	fmt.Println("15 / 3 = ", answer4)
	if answer4 != 5 {
		fmt.Println("error")
		return
	}

	fmt.Println("clear!!")
}

// 足し算 この中を書き換える
func add(x int, y int) int {
	return x
}

// 引き算 この中を書き換える
func sub(x int, y int) int {
	return x
}

// 掛け算 この中を書き換える
func mul(x int, y int) int {
	return x
}

// 割り算 この中を書き換える
func div(x int, y int) int {
	return x
}
