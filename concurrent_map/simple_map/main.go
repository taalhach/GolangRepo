package main

import (
	"sync"
	"time"
	"fmt"
)
type BTCBal struct {
	sync.RWMutex
	Balances map[string]float64
}
func main()  {
	vals:=BTCBal{
		Balances:map[string]float64{},
	}
	go vals.Adder("123",1)
	go vals.Adder("123",1)
	time.Sleep(time.Millisecond*500)
	fmt.Println(vals.Balances)

}

func (bal *BTCBal)Adder(userID string, balance float64)  {
	x:=time.Now()
	bal.Lock()
	bal.Balances[userID]+=balance
	bal.Unlock()
	fmt.Println(time.Since(x))

}