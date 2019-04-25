package main

import (
	"crypto/rand"
	"math/big"
	"log"
	"fmt"
)

func main() {
	n,err:=rand.Int(rand.Reader,big.NewInt(10))
	if err!=nil{
		log.Println("Error ")
		return
	}
	c:=n.Int64()
	fmt.Printf("%T %v %T %v ",n,n,c,c)
}