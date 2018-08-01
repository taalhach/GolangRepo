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
	go sum(x[:len(x)/2],c)
	go sum(x[len(x)/2:],c)
	fmt.Println("sum = ",<-c,<-c)
}
