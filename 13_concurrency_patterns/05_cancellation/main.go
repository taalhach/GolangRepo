package main

import (
	"fmt"
	"time"
	"sync"
)
func Write(wg *sync.WaitGroup)  {
	time.Sleep(time.Second*5)
	wg.Done()

}
func main() {
	c:=make(chan int)
	done :=make(chan struct{})
	go WriteTo(done,c)
	go ReadFrom(c)
	time.Sleep(time.Nanosecond*10000)
	close(done)
	time.Sleep(time.Millisecond*100)

}
func WriteTo(done chan struct{},c chan int  )  {
	i:=1
	for i<5{
	select {
	case c<-i:
	case <-done:
	fmt.Print("aborting")
	return
	}
	i++
	}
	close(c)

}
func ReadFrom (c chan int)  {
	for a:=range c {
		fmt.Println(a," ")

	}
}