package main

import "fmt"

func main() {
	var name string
	fmt.Print("enter your name :")
	fmt.Scan(&name)
	fmt.Println("hello", name, "!")
}
