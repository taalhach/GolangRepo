package main

import "fmt"

func main() {
	data:=make(map[string]int)
	data["talha"]=22
	data["ali"]=23
	fmt.Println(data)
	// this will give zero value of int
	fmt.Println(data["asad"])
	fmt.Println("a method to check entry exists or not")
	val,status:=data["asad"]
	if !status{
		fmt.Println("required value doesn't exist status: ",status )
	} else {
		fmt.Println(val,status)
	}
}
