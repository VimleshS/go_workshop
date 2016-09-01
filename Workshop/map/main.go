package main

import (
	"fmt"
)

func main() {

	hash := map[int]string{}
	hash[1] = "one"
	hash[2] = "two"
	hash[1] = "ononeo"

	//To ckeck key exists
	if v, ok := hash[1]; ok {
		fmt.Printf("Key found %v\n", v)
		delete(hash, 1)
	} else {
		fmt.Printf("Not Key found \n")
	}

	print_map(hash)
}

func print_map(h map[int]string) {
	fmt.Println("--------------------------------------")
	for k, v := range h {
		fmt.Printf("key %d Value %s\n", k, v)
	}

}
