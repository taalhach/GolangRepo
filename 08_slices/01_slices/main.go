package main

import "fmt"

var slice =make([]int ,5,10)


func main()  {
	slice[0]=5
	slice[1]=6
	slice[2]=7
	slice[3]=8

	//fmt.Println(slice)
	////fmt.Println(slice[6])
	//for i,_:=range slice{
	//	slice[i]=i+10
	//}
	//fmt.Println(slice)
	//slice=append(slice,60)
	//fmt.Println(slice," Capacity: ",cap(slice))
	//an_slice:=[]int{1000,10004,1225}
	//slice=append(slice,an_slice...)
	//fmt.Println(slice," Capacity: ",cap(slice))
	//copy(an_slice,slice)
	//fmt.Println(an_slice)
	//byteslice:=make([]byte,5)
	//copy(byteslice,"Hey there")
	//fmt.Println("String in bytes" ,byteslice)
	//fmt.Println(slice)
	//dequeue(slice)
	//fmt.Println("after ",slice)
	//slc:=[]int{1,2,3,4,5,6}
	//
	//q:=OrderQueue{slc}
	//q.Enqueue(5)
	//fmt.Println(q)
	//q.Dequeue()
	//q.Dequeue()
	//fmt.Println(q.OrderList)
	//fmt.Println(slc)
	a:=str{"talha",15}
	b:=a
	b.val=5
	UpdateStruct(&a,10 )
	fmt.Println("a",a)
	fmt.Println("b",b)
	for _,j:=range slice{
		slice=slice[1:]
		fmt.Println(j,slice)
	}







}
type str struct {
	name string
	val int
}
func UpdateStruct( a *str ,val int)  {
	a.val=val
}
type OrderQueue struct {
	OrderList []int
}

func (q *OrderQueue)Enqueue(order int)  {
	q.OrderList=append(q.OrderList,order)
}
func (q *OrderQueue)Dequeue()int  {
	val:=q.OrderList[0]
	q.OrderList=q.OrderList[1:]
	return val
}
func dequeue(sl []int)  {
	sl=sl[1:]
}