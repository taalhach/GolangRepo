package main

import "fmt"

func main()  {
	//func have no name are anonymous funcs normally anonymous funcs are assigned to variables
	//and called through those variables
	// following func is self executing anonymous function

	func() {
		str:="This variable is declared and called in self executing anonymous function "
		fmt.Print(str)
	}()  //these brackets lets compiler execute that func
}
