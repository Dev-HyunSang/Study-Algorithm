package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	scope, _ := strconv.Atoi(scanner.Text())

	if scope >= 90 {
		fmt.Println("A")
	} else if scope >= 80 {
		fmt.Println("B")
	} else if scope >= 70 {
		fmt.Println("C")
	} else if scope >= 60 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
