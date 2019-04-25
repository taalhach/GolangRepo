package main

import (
	"github.com/taalhach/01_GolangRepo/08_slices/intSort/int64arr"
	"fmt"
	"sort"
)

var c =make(int64arr.Int64arr,0,12)

func main() {
	j:=0
	c=int64arr.Int64arr{0,4,2,1,6,3,5}
	sort.Sort(c)
	fmt.Println(c)
	for i:=0;i<c.Len();i++{
		fmt.Println(i)
		if i==1{
			c=c[i+1:]
			break
		}
		j++
	}
	fmt.Println(c)
}


