package main

import "fmt"

func main()  {
	fmt.Println("This programm prints odd and event numbers from a collection")
	odd,even:=getOddEven(0,50);
	fmt.Println("Even numbers between 1 to 50: ",even)
	fmt.Println("Odd numbers between 1 to 50: ",odd)
}
func getOddEven(i int,j int) ([]int ,[]int)  {
	var odd []int
	var even []int
	for i<=j{
		if(i%2==0){
			even= append(even,i)
		}else {
			odd=append(odd,i)
		}
		i++
	}
	return odd,even
}
