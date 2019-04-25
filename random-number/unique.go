package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	for i:=0;i<10;i++{
		c:=rand.Int63n(10)
		fmt.Println(c)
	}
}
