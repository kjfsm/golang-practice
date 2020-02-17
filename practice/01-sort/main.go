package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 2, 6, 4, 0} //aが配列
	answer := sort(a)               //
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
	// 変えた！
	for i := 0; i < 10; i++ {
		for j := 0; j < 6; j++ {

			mae := a[j]
			ato := a[j+1]

			if mae > ato {
				a[j] = ato
				a[j+1] = mae
			}
		}
	}
	//ここまで！
	return a
}
