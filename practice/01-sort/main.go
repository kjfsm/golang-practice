package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 2, 6, 4, 0}
	answer := sort(a)
	fmt.Println(answer)

	for index, value := range answer {
		if index != value {
			fmt.Println("error")
			fmt.Println("index", index, "value", value)
			return
		}
	}
	fmt.Println("clear!!")
}

func sort(a []int) []int {
	// ここを変える
	return nil
}
