package main

import "fmt"

func main()  {
	// rune is literal in which every char is represented by a number that is ASCII code
	chr :=rune('a')
	fmt.Println("ASCI code",chr)
	fmt.Println(string(chr))

	}