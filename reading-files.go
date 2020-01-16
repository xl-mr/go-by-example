package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error)  {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "./dat"
	dat, err := ioutil.ReadFile(filename)
	check(err)

	fmt.Print(string(dat))

	f, err := os.Open(filename)
	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	defer f.Close()
}
