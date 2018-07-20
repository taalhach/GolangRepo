package main

import (
	"sync"
	"fmt"
	"time"
)
var wg sync.WaitGroup

type queue struct {
	orders []int
	sync.RWMutex
}
var m=make(map[int]*queue)
var i  =5
func main() {
	q:=queue{
		orders:[]int{0,1,2,3,4},
	}
	m[1]=&q
	a:=m[1]
	fmt.Println(m[1])

	a.delete()
	go a.delete()

	time.Sleep(time.Millisecond*9000)
	wg.Wait()
	fmt.Println(m[1])
}
func (q *queue)delete()  {
	fmt.Println(time.Now())
	time.Sleep(time.Millisecond*2000)
	fmt.Println("called ",i-1)
	q.Lock()
	q.orders=q.orders[:i-1]
	q.Unlock()
	i--
	fmt.Println(m[1])



}