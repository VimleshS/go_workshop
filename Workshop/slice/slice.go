package main

import (
	"fmt"
	"net/smtp"
)

func main() {
	slice := []string{"One", "Two", "Three"}
	//	print_slice(slice)

	//Append to slice
	slice = append(slice, "four", "fix", "six")
	//	print_slice(slice)

	//Delete from slice
	slice = append(slice[:2], slice[3:]...)
	print_slice(slice)
}

func print_slice(slice []string) {
	fmt.Println("-----------------------------------")
	//WAY 1
	/*
		for i := 0; i < len(slice); i++ {
			fmt.Printf("index %d value %s\n", i, slice[i])
		}
	*/

	//WAY -2
	for i, v := range slice {
		fmt.Printf("index %d value %s\n", i, v)
	}
}
