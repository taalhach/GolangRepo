package main

import "fmt"

func try_to_change(value string)  {

	fmt.Print("Adress of value in try_to_change func: ")
	fmt.Printf("%p\n", &value)
	value="No one"
	fmt.Println("value in try_to_change: ",value)
}
func main()  {
	value:="Who are you";  // given name

	fmt.Print("Adress of value in main func: ")
	fmt.Printf("%p\n",&value)
	try_to_change(value)
	fmt.Println("value after change process: ",value) //value won't change
	fmt.Println("Note: value is not changed")

}