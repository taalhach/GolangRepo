package main

import (
	"time"
	"fmt"
)

type S map[string]chan int
type R map[string ] S
var m R
func main()  {


	m["abc"]["abc"]<-5
	m["abc"]["abc"]<-5
	m["abc"]["abc"]<-5
	time.Sleep(1*time.Millisecond)
	m["abc"]["abc"]<-2
	time.Sleep(1*time.Second)

}

func init(){
	m=make(R)
	d:=make(S)

	d["abc"]=make(chan int)

	m["abc"]=d
	for _,v:=range m{
		for _,c:=range v{
			go func() {
				for d:=range c{
					fmt.Println("dumping: ",d)
				}
			}()
		}
	}
}


