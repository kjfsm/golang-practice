package main

import "fmt"

func main() {
	values := map[int]int{
		8:  40320,
		1:  1,
		4:  24,
		6:  720,
		0:  1,
		12: 479001600,
	}
	//key = iで　valueの中身回まわす
	for key, value := range values {
		answer := factorial(key) //Keyのかいじょうをつくる
		fmt.Println(answer)
		if answer != value {
			fmt.Println("error")
			return
		}
	}
	fmt.Println("clear!!")
}

//
func factorial(a int) int {
	// ここを変える
	res := 1
	for i := a; i > 0; i-- {

		res = res * i
	}
	a = res
	//変えたー
	return a
} //
