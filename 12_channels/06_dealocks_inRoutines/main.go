package main

import "fmt"

func main() {
	ch:=make(chan string)
	//uncomment next line to get error
	//ch<-"hello"    // this line write on channel but there is no goroutine that is receiving from the channel so deadlock occure
	go Unloader(ch)
	ch<-"hello"

	fmt.Println(<-ch,"\n","done")
}
func Unloader(ch chan string)  {
	fmt.Print(<-ch)
	//uncomment next line to get deadlock error
	ch<-" there"
}

/*Reason of dead lock

One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel,
then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program
will panic at runtime with Deadlock.

Similarly if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write
data on that channel, else the program will panic.


*/