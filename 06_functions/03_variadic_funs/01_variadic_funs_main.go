package main

import "fmt"

func main()  {

	fmt.Println("Average of nums",getAverage(1,25,35,6,6,3))
	
}
func getAverage(param ...int ) float64  {
	totalSum:=0
	fmt.Print("type of passed parameters is slice of int as given: ")
	fmt.Printf("%T\n",param)
	for _,val:=range param  {
		totalSum +=val

	}
	return float64(totalSum)/float64(len(param))

}
