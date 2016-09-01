package main

import (
	"errors"
	"fmt"
)

func main() {
	//	fmt.Println("Program starts")
	res, err := calculate(5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

func calculate(value int) (int, error) {
	if value < 10 {
		return 0, errors.New("Value less than ten")
	}
	return value * 2, nil
}
