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
func FanOut(c <-chan int) (<-chan int,<-chan int)   {
	ch:=make(chan int)
	d:=make(chan int)
	var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	for val:=range c {
	//		ch<-val
	//	}
	//	close(ch)
	//	wg.Done()
	//}()
	//wg.Add(1)
	//go func() {
	//	for val:=range c {
	//		d<-val
	//	}
	//	close(d)
	//	wg.Done()
	//}()
	wg.Add(1)
	go WriteToChan(c,d,&wg)
	wg.Add(1)
	go WriteToChan(c,ch,&wg)


	go func() {
		wg.Wait()
	}()
	return d,ch
}
func WriteToChan(s <-chan int,d chan <-int,wg *sync.WaitGroup)  {
	defer wg.Done()
	for val:=range s{
		d<-val
	}
	close(d)
}

func FanIn(cList ...<-chan int ) <-chan int {
	ch:=make(chan int)
	var wg sync.WaitGroup
	writer:= func(i <-chan int) {
		for val:=range i{
			ch<- val*2
		}
		wg.Done()
	}
	wg.Add(len(cList))
	for _,c:=range cList{
		go writer(c)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func Print(c <-chan int)  {

	for val:=range c{
		fmt.Println("Printer prints: ",val)
	}
}

func main()  {

	Print(FanIn(FanOut(Generator(1,2,3,4,5,6))))

}