package main

import "fmt"

func Genrate_numbers(num ...int) chan int {
	c:=make(chan int)

	go func() {

		for _,val:=range num{
			c<-val
		}
		close(c)
	}()
	return c
}

func SquareGenerator(nums <-chan int ) chan int{
	c:=make(chan int)
	go func() {
		for val:=range nums{
			c<-val*val
		}
		close(c)
	}()
	return c
}

func Print(sq <-chan int)  {
	for vals:=range sq{
		fmt.Println(vals)

	}
}

func main() {
	Print(SquareGenerator(Genrate_numbers(2,3,4,5,6)))
}
