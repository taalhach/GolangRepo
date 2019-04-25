package main

import "sync"

type queue struct {
	sync.RWMutex
	Numbers []int
}

func (q *queue)()  {
	
}

func main()  {
	a:=queue{
		Numbers:[]int{1,2,3,4},
	}
	
}