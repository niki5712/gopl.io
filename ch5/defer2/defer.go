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
	defer func() {
		i := 1
		if p := recover(); p != nil {
			fmt.Println(fmt.Errorf("error instead of panic: %v, %v, %v", p, i, x))
		} else {
			defer func() {
				if p := recover(); p != nil {
					fmt.Println(fmt.Errorf("new error instead of new panic: %v, %v, %v", p, i, x))
				}
			}()
			i = 0 / (i - 1)
		}
	}()
	defer fmt.Printf("defer %d\n", x)
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	f(x - 1)
}
