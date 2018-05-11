package main

import "fmt"

func main()  {
	slice:=make([]string, 5 , 10)
	for i,_:=range slice{
		slice[i]=string(65+i)
	}
	fmt.Println(slice)
	fmt.Println("Slice: ",slice)
	slice=append(slice[0:len(slice)-1])   // this append is deleting an element from last of slice
	fmt.Println(slice)

}

