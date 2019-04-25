package main

import (
	"sort"
	"fmt"
)
type Int64arr []int64

func (a Int64arr) Len() int           { return len(a) }
func (a Int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Int64arr) Less(i, j int) bool { return a[i] < a[j] }
func (a Int64arr) String() (s string) {
	sep := "" // for printing separating commas
	for _, el := range a {
		s += sep
		sep = ", "
		s += fmt.Sprintf("%d", el)
	}
	return
}

var a Int64arr

func main() {
	a=[]int64{3,1,6,90,34,3}
	sort.Sort(a)
	sort.Sort(sort.Reverse(a))
	fmt.Println(a)
}


