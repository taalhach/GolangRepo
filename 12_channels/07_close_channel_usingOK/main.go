package main

import (
	"fmt"
	"time"
)

func main()  {
	ch:=make(chan int)
	go Producer(ch)
	go consumer(ch)
	time.Sleep(time.Millisecond*5000)
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
	for{
		 val,ok:=<-ch
		if !ok{
			fmt.Println("channel closed and zero value of channel is: ",val)
			break
		}
		fmt.Println(val)
	}
}