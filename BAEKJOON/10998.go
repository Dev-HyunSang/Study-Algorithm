package main

import "fmt"

var (
	a int
	b int
)

func main() {
	fmt.Scanf("%d %d", &a, &b)
	fmt.Print(a * b)
}
