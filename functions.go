package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3))
}

func add(a int, b int, c int) int {
	return a + b + c
}
