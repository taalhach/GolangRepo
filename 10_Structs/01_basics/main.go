package main

import "fmt"

type car struct {
	name string
	model string
	powerInCC float32
	year int32
}
func main() {
	mycar:=car{model:"swept tail",
	name:"Rolls Royce",
	powerInCC:33.5,
	year:2018,
	}
	fmt.Println(mycar)
	mycar2:=new(car)  // new is in notes
	fmt.Println(mycar2)   //prints pointer to zero values
	fmt.Println(&mycar2)   //prints address of mycar2

}
