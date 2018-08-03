package main

import (
	"fmt"
	"time"
)

func sum(x []int,c chan int) {
	fmt.Println(time.Now())
	sum:=0
	for _,val:=range x{
		sum+=val
	}
	fmt.Println(time.Now()," write")
	c<-sum
}
func main()  {
	x:=[]int{ 1,2,4,5,6,7}
	c:=make(chan int)
	ch:=make(chan string)
	go write(ch)
	go sum(x[:len(x)/2],c)
	go sum(x[len(x)/2:],c)


	go Reflect(ch)
	fmt.Println("sum = ",<-c,<-c,<-ch)
	time.Sleep(time.Millisecond*5000)

	}
func write(c chan string)  {
	c<-"Hey there"
	c<-"Hey there 22"

}
func Reflect(c chan string){
	time.Sleep(time.Millisecond*5000)
	x:=<-c
	fmt.Println(x," In Reflect ")
}
