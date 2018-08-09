package type2

import "fmt"

type Order interface {
	BuyOrder()[]Order
	SellOrder()[]Order
}

type Str struct {
	Id string
}

func (p Str)BuyOrder() []Order {

	c:=[]Order{
		p,
	}
	return c
}
func (p Str)SellOrder() []Order {
	c:=[]Order{
		p,
	}
	return c
}

func PrintW(i Order)  {
	fmt.Println(i)
}

