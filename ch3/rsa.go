package main

import (
	"fmt"
	"math"
)

//扩展欧几里德算法
func exgcd(a int, b int) (int, int) {
	if b == 0 {
		return 1, 0
	}
	x, y := exgcd(b, a%b)
	x, y = y, x-(a/b)*y
	return x, y
}

func rsa() (int, int, int) {
	//第一步，随机选择两个不相等的质数p和q。
	p := 3
	q := 11
	//第二步，计算p和q的乘积n。
	n := p * q
	//第三步，计算n的欧拉函数φ(n)。
	phi := (p - 1) * (q - 1)
	//第四步，随机选择一个整数e，条件是1< e < φ(n)，且e与φ(n) 互质。
	e := 3
	//第五步，计算e对于φ(n)的模反元素d。
	d, _ := exgcd(e, phi)
	if d < 0 {
		d = phi + d
	}
	//第六步，将n和e封装成公钥，n和d封装成私钥。
	return n, e, d
}

func crypt(msg int, key int, length int) int {
	return int(math.Pow(float64(msg), float64(key))) % length
}
func main() {
	n, e, d := rsa()
	fmt.Printf("public key = %d,private key = %d,length = %d\n", e, d, n)

	m := 24
	fmt.Println("Origin Message = ", m)

	encodeMsg := crypt(m, e, n)
	fmt.Println("Encoded Message = ", encodeMsg)

	decodeMsg := crypt(encodeMsg, d, n)
	fmt.Println("Decoded Message = ", decodeMsg)
}
