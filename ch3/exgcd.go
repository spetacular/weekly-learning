package main

import (
	"fmt"
)

//扩展欧几里德算法
func exgcd(a int, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	x, y := exgcd(b, a%b)
	x, y = y, x-(a/b)*y
	fmt.Printf("%d\t%d\t%d\t%d\n", a, b, x, y)
	return x, y
}
func main() {
	a := 43
	b := 17
	fmt.Printf("a\tb\tx\ty\n")
	x, y := exgcd(a, b)
	fmt.Printf("Result is : x=%d , y=%d\n", x, y)
}
