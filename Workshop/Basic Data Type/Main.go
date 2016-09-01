package main

import (
	"fmt"
)

func main() {
	fmt.Println(CtoF(30))
	fmt.Println(FtoC(86))
}

func CtoF(celcius float32) float32 {
	return celcius*1.8 + 32
}

func FtoC(fahrenheit float32) float32 {
	return (fahrenheit - 32) / 1.8
}
