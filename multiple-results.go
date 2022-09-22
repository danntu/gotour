package main

import "fmt"

func main() {

	a, b := swap(9, 81)
	fmt.Println(swap(a, b))
	fmt.Println(swap(87, 98))
}

func swap(i, j int) (int, int) {
	return j, i
}
