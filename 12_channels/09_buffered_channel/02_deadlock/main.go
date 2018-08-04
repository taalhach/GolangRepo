package main

import "fmt"

func main() {
	ch :=make(chan interface{},2)
	fmt.Println("first write")
	ch<-1
	fmt.Println("Second write")
	ch<-2
	go Read(ch)
	ch<-3
	fmt.Println("Done third write")
	fmt.Println("4th write panics program bcz no concurrent read from channel is being performed this is deadlock")
	ch<-4


}
func Read(ch <-chan interface{})  {
	fmt.Println("Read done: ",<-ch)
}
