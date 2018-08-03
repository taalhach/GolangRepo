package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	 inte :=make(chan int )
	 sq:=make(chan int )
	go Generator(inte)
	go Squarer(inte,sq)
	go Printer(sq)
	time.Sleep(time.Millisecond*500)
}
func Generator(c chan<- int)  {
	c<-2
}
func Squarer(c <-chan int,sq chan <-int)  {
	x,ok:=<-c
	fmt.Println("Read: ",x)
	if !ok{
		log.Println("Empty channel")
	}
	// un comment next line to see error
	//c<-12
	sq<-x*x

}
func Printer(c <-chan int)  {
	a,ok:=<-c
	if !ok{
		log.Println("Empty channel")
	}
	fmt.Println(a)
}