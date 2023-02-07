package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("%s", name)
}
