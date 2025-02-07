package main

import "fmt"

func main() {
	n := 5
	if n > 0 {
		fmt.Println("5 is positive")
	} else if n < 0 {
		fmt.Println("5 is negative")
	} else {
		fmt.Println("zero")
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
