package main

import (
	"fmt"
	"time"
)

func WriteToChannel(ch chan<- int)  {
	i:=0
	for i<5{
		ch<-i+1
		fmt.Println("Iteration no: ",i)
		i++
	}
	close(ch)
}
func ReadFromChannel(ch <-chan int)  {
	for val:=range ch{
		fmt.Println("Read: ",val," I am going to sleep")
		time.Sleep(time.Millisecond*500)
	}
}
func main() {
	ch:=make(chan int,2)
	fmt.Println(cap(ch))
	go WriteToChannel(ch)
	go ReadFromChannel(ch)
	time.Sleep(time.Millisecond*5000)

}

