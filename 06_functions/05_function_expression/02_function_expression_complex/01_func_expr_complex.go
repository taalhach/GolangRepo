package main

import "fmt"

// this programs divides a number by 7 and gives remainder

func getRemainder(input int) func() int  {
	return func() int {
		return input%7
	}

}
func main()  {
	value:=getRemainder(13)
	fmt.Println("Remainder is: ",value())
	fmt.Printf("%T\n",value)
}
