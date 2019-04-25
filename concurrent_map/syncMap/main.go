package main

import (
	"sync"
	"fmt"
	"time"
)

type Balances struct {
	B sync.Map
}

var BTCBalance =Balances{}

func (bal *Balances)UPdater(userID string, balance float64)  {
	x:=time.Now()
	v,_:=bal.B.Load(userID)
	if v==nil{
		bal.B.Store(userID,balance)
		fmt.Println("nah")
		fmt.Println(time.Since(x))
		return
	}
	b,_:=v.(float64)
	b+=balance
	bal.B.Store(userID,b)
	fmt.Println(time.Since(x))
}



func main() {
	go BTCBalance.UPdater("123",1.23)
	go BTCBalance.UPdater("123",1.23)
		time.Sleep(time.Millisecond*1500)
	fmt.Println(BTCBalance.B.Load("123"))
}