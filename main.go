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


		res2:=calc.Div(60,20)
	fmt.Println(res2)


	res1 := calc.Multi(10, 20)
	fmt.Println(res1)
}
