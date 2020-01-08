package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

type Engine struct {
	Start int
	Stop int
}

type Car struct {
	Engine        //包含 Engine 类型的匿名字段
}

func (c *Car) work()  {
	fmt.Println(c.Start + c.Stop)
}

func main() {
	go say("world")
	say("hello")

	var a string = "aaa"

	fmt.Println("aaa" + a + "bbb")
}
