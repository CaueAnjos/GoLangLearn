package main

import "fmt"

func main() {
	x := 10
	if true {
		x := 5
		fmt.Println(x) // 5
	}
	fmt.Println(x) // 10
}
