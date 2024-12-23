package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Print(i, " ")
		}
	}

	fmt.Print("\n")

	if x%2 == 0 || x%3 == 0 || x%5 == 0 || x%7 == 0 && x != 1 && x != 2 && x != 3 && x != 5 && x != 7 {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}
}
