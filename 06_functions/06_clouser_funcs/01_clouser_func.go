package main

import "fmt"
// enclosing a variable scope into a function that is inside the other func so that we can access outer func variables
//clouser enables a user to access function level variables into other functions note that these variables are not declared on file level

func main()  {
	input:="hey there"
	getCount:=func () string{
		input="changed"
		return input
	}
	fmt.Print(getCount())
}

