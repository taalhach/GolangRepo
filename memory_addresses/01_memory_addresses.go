package main

import "fmt"
var yardCon =1.09;
func main()  {
	meter:=0.0
	fmt.Print("Enter number:")
	fmt.Scan(&meter)
	fmt.Println("address:  ",&meter)
	fmt.Printf("%d\n",&meter)
	fmt.Println("yards:  ",meter*yardCon)
	}
