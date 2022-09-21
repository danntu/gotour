package main

import "fmt"

func main() {
	fmt.Println(split(22))
}

func split(sum int) (x, y int) {
	x = sum / 3
	y = sum - x
	return
}
