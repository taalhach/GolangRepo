package main

import (
	"sort"
	"fmt"
)

func main()  {
	s:=[]int{2,6,4,0,5}
	a:=sort.IntSlice(s)
	sort.Sort(a)
	fmt.Println("Ascending order: it ",s)
	sort.Sort(sort.Reverse(a))
	fmt.Println("Descending order: ",s)
}
