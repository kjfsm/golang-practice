package main

import "fmt"

type hashmap struct {
	value []data
}

type data struct {
	key   string
	value string
}

func main() {
	hashmap := hashmap{}
	hashmap.insert("key1", "value1")
	value, error := hashmap.get("key1")
	if error != nil {
		fmt.Println(error)
		return
	}

	if value != "value1" {
		fmt.Println("error value1", value)
	}
}

func (hashmap *hashmap) insert(key string, value string) {

}

func (hashmap *hashmap) get(key string) (string, error) {
	return "", nil
}
