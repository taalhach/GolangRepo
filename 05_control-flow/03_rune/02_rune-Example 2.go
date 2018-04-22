package main

import "fmt"

func main()  {
	for i:=65;i<500;i++{
		fmt.Println(string(i),[]byte(string(i)))
	}
}
