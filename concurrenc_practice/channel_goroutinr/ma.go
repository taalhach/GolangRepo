package main

import (
	"time"
	"fmt"
)

var ch chan int
func main()  {
	wraper()
}
func wraper()  {
	go writer(1)
	go writer(1)
	go writer(1)
	go writer(1)
	go writer(1)
	go writer(1)
	for  {
		go writer(1)
		time.Sleep(1*time.Second)
	}
}

func callme(v int )  {
	fmt.Print("read ",v)
	time.Sleep(1*time.Second)
	fmt.Println("db ",v)

}
func writer(v int)  {
	ch<-v
}

func init() {
	ch=make(chan int)
	go func() {
		for c:=range ch{
			go callme(c)
		}
	}()
}