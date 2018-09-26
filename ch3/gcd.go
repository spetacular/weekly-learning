package main

import (
	"fmt"
)

//辗转相除法求最大公约数
func gcd(a int, b int) int {
	mod := a % b
	fmt.Printf("%d\t%d\t%d\n", a, b, mod)
	if mod == 0 {
		return b
	}
	return gcd(b, mod)
}
func main() {
	a := 1251
	b := 390
	fmt.Printf("a\tb\tmod\n")
	c := gcd(a, b)
	fmt.Printf("The highest common divissor between %d and %d is : %d\n", a, b, c)
}
