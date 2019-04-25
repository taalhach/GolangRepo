package main

import (
	"math/big"
	"fmt"
)

func main() {
	a:=big.NewInt(1000)
	fmt.Printf("%T",a.Sub(a,big.NewInt(5)).Mod(a,big.NewInt(2)))
}
