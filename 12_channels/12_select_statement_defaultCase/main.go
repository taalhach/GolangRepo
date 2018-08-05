package main
import (
	"fmt"
)

var balanes =map[string]float64{
	"abc":0.333,
	"a":1,
	"b":2,
}

func BalancerReponse(Uid string,ch chan<- float64)  {
	ch<-balanes[Uid]
}

func DatabaseReponse(Uid string,ch chan<- float64)  {
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
		fmt.Printf("User %v has %v balance \n Response from balancer",userId,b)
	case b := <-db:
		fmt.Printf("User %v has %v balance \n Response from DB",userId,b)
	}
}
