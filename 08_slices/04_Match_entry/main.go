package main

import (
	"sort"
	"fmt"
	"time"
)

var v  =[]int {2,1,8,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,322,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,322,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32,2,1,2,3,2,5,7,6,3,1,2,3,4,5,6,7,8,7,55,4,3,2,23,32,23,2,32,3,2,32,3,2,32,32,2,32,2,3,2,3,2,32,32,2,32,32}
func main()  {
	x:=time.Now()
	t:=time.Since(x)
	sort.Ints(v)
	fmt.Println(t,len(v))

	y:=GetMatch(9)
	fmt.Println(y)
}

func GetMatch(mp int) []int  {
	x:=time.Now()

	p:=[]int{}
	for _,val:=range v{
		if val>mp{
			break
		}
		p=append(p,val)
	}
	t:=time.Since(x)
	fmt.Println(t)
	return p
}


