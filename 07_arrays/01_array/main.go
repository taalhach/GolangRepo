package main

import "fmt"

func main()  {
var array [5]int

for i,_:= range array{
fmt.Println(string(i+65))
}
}