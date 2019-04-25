package main

import "fmt"

type ar []int
var array ar
func main() {
	array=[]int{1,2,3,4,3}
	array.checkITouT(1)
	fmt.Println(array)
}
func ( arr *ar)checkITouT(amount int)  {
	for i,a:=range *arr{
		if amount==a{
			array=array[i+1:]
			break
		}else if a<amount {
			amount-=a
			continue

		}else if a>amount {
			a-=amount
			array[i]=a
			array=array[i:]
			break
		}
	}
}
