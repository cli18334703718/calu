package main

import (
	"fmt"
	"gittest/calc"
)

func main() {

	res := calc.Add(10, 20)
	fmt.Println(res)

	result := calc.Sub(30, 10)
	fmt.Println(result)

	res1 := calc.Multi(10, 20)
	fmt.Println(res1)
}
