package main

import (
	"fmt"
	"time"
)

var balanes =map[string]float64{
	"abc":0.333,
	"a":1,
	"b":2,
}

func BalancerReponse(Uid string,ch chan<- float64)  {
	time.Sleep(time.Millisecond*1000)
	ch<-balanes[Uid]
}

func DatabaseReponse(Uid string,ch chan<- float64)  {
	time.Sleep(time.Millisecond*6000)
	fmt.Println("Conecting db")
	fmt.Println("Finding result")
	time.Sleep(time.Millisecond*1000)
	ch<-balanes[Uid]
}

func main() {
	blncr:=make(chan float64)
	db:=make(chan float64)
	userId:="abc"
	go DatabaseReponse(userId,db)
	go BalancerReponse(userId,blncr)
	select {
	case b:=<-blncr:
		fmt.Printf("User %v has %v balance",userId,b)
	case b := <-db:
		fmt.Printf("User %v has %v balance",userId,b)
	default:
		fmt.Println("Couldn't wait that much long ")
	}

}
