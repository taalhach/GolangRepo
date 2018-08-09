package main

import (
	"fmt"
	"github.com/taalhach/01_GolangRepo/09_maps/06_type_map/types"
	"github.com/taalhach/01_GolangRepo/09_maps/06_type_map/type2"
)

func main() {
	var l types.Orderbook=types.Orderbook{}
	fmt.Println(l)
	l[1.23]=[]int{1,2,3}
	var i =type2.Str{
		Id:"1aaa",
	}
	type2.PrintW(i)

}