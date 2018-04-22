package main

import "fmt"

func main()  {
	// represented in backticks contains spaces new lines and double quotes
	raw_str:=`Hello there this is a "raw" string
which can contain lines, spaces single and double quotes like this '' "" `
	fmt.Println(raw_str)
}
