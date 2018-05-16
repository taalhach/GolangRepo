package main
import "fmt"
type amap map[string]int
var a amap
func main()  {
	fmt.Println(a)
	a["1"]=15
	fmt.Println(a) // entry to nil map error
}