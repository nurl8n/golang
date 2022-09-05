package main

import "fmt"

func main() {
	fmt.Println("hello world")
	// hello world

	fmt.Println("go" + "lang")
	// golang

	fmt.Println("1+1 =", 1+1)
	// 1+1 = 2

	fmt.Println("7.0/3.0 =", 7.0/3.0)
	// 7.0/3.0 = 2.333333333

	fmt.Println(true && false)
	// false

	fmt.Println(true || false)
	// true

	fmt.Println(!true)
	// false

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	/*
		1
		2
		3
	*/
}
