package main

import "fmt"

type orderbook map[string]int
func main() {
	fmt.Println("nil entry error will occure so use make func to initialize map")
	var ob orderbook
	ob["1.1"] = 2
	fmt.Print(ob)


	}
