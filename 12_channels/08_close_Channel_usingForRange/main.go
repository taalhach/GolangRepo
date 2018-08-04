package main

import (
	"fmt"
	"time"
)

func main()  {
	ch:=make(chan int)
	go Producer(ch)
	go consumer(ch)
	time.Sleep(time.Millisecond*1/10)
}
func Producer(ch chan int)  {
	i:=0
	for i<10{
		ch<-i
		i++
	}
	close(ch)
	fmt.Println("Done")
}
func consumer(ch chan int)  {
	for val:=range ch{
		fmt.Println(val)
	}
}