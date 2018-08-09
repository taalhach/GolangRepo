package main

import (
	"github.com/taalhach/01_GolangRepo/11_intreface/masla/types"
	"sort"
	"fmt"
)

func main()  {
	s:=&types.Str{
		Name:"talha",
	}
	types.CallIn(s)


	sb:=sort.Float64Slice{}
	i:=0.0

	for i<5{
		sb = append(sb,1.0)
		i++
	}
	fmt.Println(sb)
	for _,i=range sb{
		fmt.Println(i)
	}

}

