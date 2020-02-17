package main

import "fmt"

func main() {
	i := 42
	
	p1 := &i        // point to i 参照渡し
	fmt.Println("p1",p1) //アドレス
	fmt.Println("*p1",*p1) //中身 read i through the pointer
	p2 := *p1
	*p1 = 21        // set i through the pointer
	fmt.Println("p2",p2)
	fmt.Println("i",i) // see the new value of i
    
	j := 43
	k := j         // 実体渡し jが大きいものだと、処理に時間がかかる
	fmt.Println("k",k) // read i through the pointer
	k = 21        // set i through the pointer
	fmt.Println("j",j) // see the new value of i
}
