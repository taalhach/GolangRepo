package main

import (
	"time"
	"fmt"
)

var j []int
func main() {
	ch:=time.Tick(1*time.Second)
		i:=0
		for range ch{
			go wanaappend(i)
			i++
		}


}

func wanaappend(i int)  {
	if len(j)<=20{
		j=append(j,i)
	}else {
		j=j[1:]
		j=append(j,i)
	}

		fmt.Println(j)

}
func init() {
	j=make([]int,0,20)
}


