package main

import "fmt"

func main() {
	var (
		x int
		y int
	)
	fmt.Scanf("%d\n", &x)
	fmt.Scanf("%d\n", &y)
	if x > 0 {
		if y > 0 {
			fmt.Println(1)
		} else {
			fmt.Println(4)
		}
	} else {
		if y > 0 {
			fmt.Println(2)
		} else {
			fmt.Println(3)
		}
	}
}
