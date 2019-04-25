package main

import "fmt"

type Order interface {
	Havechan()
}

func CallMe(i Order)  {
	i.Havechan()
}

type engineMOrder struct {
	OrderID string
	amount int64
}
func (od engineMOrder)Havechan()  {

}

type engineLOrder struct {
	OrderID string
	amount int64
	price  int64
}

func (od engineLOrder)Havechan()  {

}

type orderQueue map[string]chan Order

type PairQueue map[string]orderQueue

func (p PairQueue)RegisterPair(pair string,orderTyes ...string)  {
	if p[pair]!=nil{
		return
	}
	orderChans:=orderQueue{}
	for _,od:=range orderTyes{
		orderChans[od]=make(chan Order)
	}
	p[pair]=orderChans
}
func (p PairQueue)UpdatePair(pair string,orderTyes ...string) {
	for _,od:=range orderTyes{
		if p[pair][od]!=nil{
			continue
		}
		p[pair][od]=make(chan Order)
	}
}



func main()  {
	m:= make(PairQueue)
	m.RegisterPair("ETH/BTC","mkt_buy","mkt_sell","lmt_buy","lmt_sell")
	m.UpdatePair("ETH/BTC","mkt_buy",)

	fmt.Printf("%T %v",m["ETH/BTC"]["lmt_buy"],m["ETH/BTC"])

}