package main

import (
	"sync"
	"fmt"
	"time"
)

type data struct {

	field string
	lock sync.Mutex
}

func main()  {
	d:=data{
		field:"abc",
	}
	go change(&d)
	d.lock.Lock()
	d.field="changing in main "
	time.Sleep(10000*time.Millisecond)
	fmt.Println(d)

}
func change(d *data)  {
	d.lock.Lock()
	time.Sleep(2000*time.Millisecond)
	d.field="some value "

	fmt.Println(*d)
}



//
//import (
//	"sync"
//	"fmt"
//	"time"
//)
//
//var mp=make(map[int]array)
//
//type array struct {
//	ordes []int
//	lock sync.Mutex
//}
//
//func main() {
//	a:=make([]int,5)
//	ar:=array{
//		ordes:a,
//	}
//	mp[1]=ar
//	ar=mp[1]
//	ar.lock.Lock()
//	go edit(mp[1])
//	time.Sleep(2000*time.Millisecond)
//	go edit2(mp[1])
//	time.Sleep(20000*time.Millisecond)
//	fmt.Println(mp)
//
//
//}
//func edit(ar array)  {
//	fmt.Println("call time: ",time.Now())
//	ar.lock.Lock()
//	time.Sleep(10000*time.Millisecond)
//	ar.ordes[0]=1
//	fmt.Println("di")
//}
//func edit2(ar array)  {
//	fmt.Println("call time 2: ",time.Now())
//	ar.ordes[2]=2
//	fmt.Println("edit2")
//
//
//}