package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	_ = os.Setenv("Foo", "1")

	fmt.Println("Foo: ", os.Getenv("Foo"))
	fmt.Println("BAR: ", os.Getenv("BAR"))

	fmt.Println()

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
