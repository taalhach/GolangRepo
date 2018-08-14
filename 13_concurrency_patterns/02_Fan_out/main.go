package main

import (
	"fmt"
	"sync"
)

func Generator(nums ...int)<-chan int  {
	c:=make(chan int)
	go func() {

		for _,v:=range nums{
			c<-v
		}
		close(c)
	}()
	return c
}
func FanOut(c <-chan int)  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for val:=range c {
			fmt.Println("func 1 reads: ",val)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for val:=range c {
			fmt.Println("func 2 reads: ",val)
		}
		wg.Done()
	}()
	wg.Wait()
}

func main()  {
	FanOut(Generator(2,3,4,5,6,7))


}