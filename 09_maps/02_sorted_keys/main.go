package main

import (
	"fmt"
	"sort"
)

var m=make(map[float64][]string)
var prices []float64
func main() {
	i:=0
	for i<20 {
		orders:=[]string{"order1","order2","order3","order4"}
		m[float64(i)+0.1]=orders
		i++
	}

	for key,val :=range m{
		fmt.Println(key," : ",val)
		prices=append(prices,key)
	}
	fmt.Println(prices)
	sort.Float64s(prices)
	fmt.Println(prices)
}
