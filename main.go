package main

import (
	"fmt"
)

func main() {
	var s string = "hello golang 学习"
	var p *string = &s
	fmt.Println("s = '", s, "'")
	fmt.Println("p = '", p, "'")
	var pp **string = &p
	fmt.Println("pp = '", pp, "'")
	**pp = "hello world"
	fmt.Println("s = '", s, "'")

}
