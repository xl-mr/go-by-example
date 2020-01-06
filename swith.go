package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write", i, "as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println(i, "is not numeric")
	}

	t := time.Now()
	switch {
	case 12 > t.Hour():
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
}
