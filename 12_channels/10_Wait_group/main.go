package main

import (
	"sync"
	"fmt"
	"time"
)

func AnotherFunc(i int,wg *sync.WaitGroup)  {
	fmt.Println("Started routine ",i)
	x:=i*1000
	time.Sleep(time.Millisecond*time.Duration(x))
	fmt.Println("Done at: ",time.Now())
	wg.Done()
}


func main() {
	var wg sync.WaitGroup
	for i:=1;i<5;i++{
		wg.Add(1)
		go AnotherFunc(i,&wg)
	}

	fmt.Println("Waiting for go routine to complete")
	wg.Wait()
	fmt.Println("Terminating!!!!!!!!")


}
