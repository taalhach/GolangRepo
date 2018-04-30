package main

import "fmt"

func main()  {
	input :=[]int{1,2,3,4,5,6,7,8,9}
	sum:=getTotalSum(input...)         // ... three dots in argument tells take values one by one and send these values as parameters in func
	fmt.Print("Total sum of ",input," is: ",sum)
	}

func getTotalSum(param ...int) int {
	totalSum:=0
	for _,val:=range param{
		totalSum+=val
	}
	return totalSum
}
