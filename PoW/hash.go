package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0x1234)
	fmt.Printf("0x%x, %v\n", a, a)
	// 设置 a 为 0x5678
	a.SetBytes([]byte{0x56, 0x78})
	fmt.Printf("0x%x, %v\n", a, a)
}
