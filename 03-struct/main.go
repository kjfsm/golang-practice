package main

import "fmt"

type pokemon struct {
	name   string
	gender string
	level  int
	moves  []string
}

func main() {
	raichu := pokemon{
		name:   "ライチュウ",
		gender: "famele",
		level:  32,
		moves:  []string{"10万ボルト", "アイアンテール", "穴を掘る", "スパーク"},
	}
	fmt.Println(raichu.name)

	p := &raichu
	fmt.Println("p", p)
	fmt.Println((*p).moves)
	fmt.Println(p.moves)
}
