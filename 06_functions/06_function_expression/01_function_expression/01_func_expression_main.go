package main

import (
	"fmt"
	"strings"
)
//assigning a variable to a function is function expression
func main()  {
	annoymous:= func(str string) string {
		return strings.ToUpper(str)
	}
	fmt.Println(annoymous("hey there"))
	fmt.Printf("%T\n",annoymous)
}