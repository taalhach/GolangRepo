package main

import (
	"github.com/taalhach/GolangRepo/Package_Example/Import_packagge"
	"fmt"
)
func main() {
	        fmt.Println("calling variables directly due to package level scope")
	fmt.Print("name: "+Import_packagge.Name+"\n username: " + Import_packagge.Username+"\n")
	fmt.Println("typical getter like other programming languages")
	fmt.Println("password: " +Import_packagge.GetPass())
	fmt.Println("getting password from file")
	fmt.Println("password: " +Import_packagge.GetPassFromFile())
	}