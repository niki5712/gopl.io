package main

import (
	"fmt"
)

func main() {
	defer printNiki()
	f(3)
	fmt.Println("2niki2")
}

func printNiki() {
	fmt.Println("!!!niki!!!")
}

func f(x int) {
	defer fmt.Printf("defer %d\n", x)
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	f(x - 1)
}
