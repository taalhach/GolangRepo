package main

import "fmt"

//func that changes value of variabl passed to it
//in this func we are passing address of variable as parameter
func try_to_Change(value *string)  {
	fmt.Println("value address: ",value) 	//prints address of value note: that this address and address of value in main is same
	*value="No one" 							// *value is dereferencing in this step we are dereferencing and assigning value
}
func main()  {
	value:="Who are you"
	fmt.Println("value address: ",&value)
	fmt.Println("value before change process: ",value)
	try_to_Change(&value)
	fmt.Println("value after change process: ",value)
	// you will notice that value will be changed to value assigned in try_to_change func
}

