package main

import (
	"fmt"
	"time"
)

type Acc map[string]float64
var b =Acc{}
func (a Acc)add(u <-chan string,b <-chan float64)  {
	id:=<-u
	if a[id]==0{
		a.Init(id)
		fmt.Println("caled init")
	}

	c:=<-b
	a[id]+=c
	fmt.Println(a)
}

func (a Acc)Init(u string)  {
	a[u]=0.0
}

func main() {
	ch:=make(chan float64)
	ch1:=make(chan string)
	go b.add(ch1,ch)
	ch1<-"user"
	ch<-1.2
	time.Sleep(time.Second*1)

	fmt.Println(b)

	fmt.Println("done")
}
