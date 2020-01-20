package main

import "time"

import "fmt"

func main() {
	go spinner(100 * time.Millisecond)

	const n = 45
	fibN := fib(n)

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	fmt.Println("验证main goroutine是main, 在main的其他go goroutine 在等main 退出就自然退出")
	fmt.Println("start check")

	time.Sleep(10 * time.Second)

	fmt.Println("end check, exit main")
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
