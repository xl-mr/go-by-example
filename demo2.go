package main

import "fmt"

func main() {
	for _, r := range "abc" {
		fmt.Printf("%c ", r)
	}
	fmt.Println()

	for _, r := range `abc` {
		fmt.Printf("%c ", r)
	}
	fmt.Println()
}
