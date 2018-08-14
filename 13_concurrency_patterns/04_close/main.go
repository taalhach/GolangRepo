package main

import (
	"sync"
	"fmt"
	"time"
)

func Read(c chan int,wg *sync.WaitGroup)  {
	for i:=1;i<16 ;i++  {
		fmt.Println(<-c)
	}
	wg.Done()

}
func Write(c chan int, wg *sync.WaitGroup)  {
	for i:=1;i<4 ;i++  {
		c<-i

	}
	close(c)
	wg.Done()

}

func main() {
	var wg sync.WaitGroup
	c:=make(chan int)
	wg.Add(2)
	go Write(c,&wg)
	wg.Add(1)
	go Read(c,&wg)
	time.Sleep(500*time.Millisecond)

}