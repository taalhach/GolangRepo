package main

import "fmt"

var m map[string]map[string]string
type c map[string]string
type p struct {
	a,b string
}
var ex map[string]p
func main()  {
	fmt.Println(m["s"],len(m))
	fmt.Println(ex["p"])
}
