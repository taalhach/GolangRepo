package main

import (
	"fmt"
)

func main()  {

	empName:="bilal"
		switch  {
	case empName=="bilal",empName=="Bilal":
		fmt.Println("Bilal arif is CEO of Block360 ")
	case empName=="Yumna",empName=="yumna":
		fmt.Println("Yumna Ghazi is co founder of Block360 ")
	case empName=="Zaryab",empName=="zaryab":
		fmt.Println("Zaryab is backend developer at Block360 ")
	case empName=="Danish",empName=="danish":
		fmt.Println("Zaryab is frontend developer at Block360 ")
	default:
		fmt.Println("Not a part of block360 ")
	}
}
