package main

import "fmt"

func Assert(i interface{})  {
	switch t:=i.(type) {
	case int:
		fmt.Println("Int value ",t/1)
	case []int:
		fmt.Println("[]int vals: ", len(t))
	}
}

func main()  {
	i:=5
	Assert(i)
	a:=[]int{
		1,2,3,4,
	}
	Assert(a)

}