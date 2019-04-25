package main

import (
	"math/big"
	"fmt"
)

func main()  {
	if big.NewInt(0).Sub(big.NewInt(10),big.NewInt(5)).Cmp(big.NewInt(3))==1{
		fmt.Println(true)
	}else {
		fmt.Println(false)
	}

}
