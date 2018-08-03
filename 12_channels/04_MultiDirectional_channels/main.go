package main

import (
	"fmt"
	"time"
	"log"
)
var inte =make(chan int )
var sq=make(chan int )
func main() {
	go Generator(inte)
	go Squarer(inte)
	go Printer(sq)
	time.Sleep(time.Millisecond*2000)
}
func Generator(c chan int)  {
	c<-2
}
func Squarer(c chan int)  {

	x,ok:=<-c
	if !ok{
		log.Println("Empty channel")
	}
	sq<-x*x


}
func Printer(c chan int)  {
	a,ok:=<-c
	if !ok{
		log.Println("Empty channel")
	}
	fmt.Println(a)
}