package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {
	point := point{1, 2}

	fmt.Println(point)

	fmt.Printf("%v\n", point)
	fmt.Printf("%+v\n", point)
	fmt.Printf("%#v\n", point)

	_ = os.Mkdir("logs", 0777)

}
