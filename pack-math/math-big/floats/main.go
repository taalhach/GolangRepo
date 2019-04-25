package main

import (
	"fmt"
	"math/big"
)

type Order struct {
	amount,price,fee float64

}
type O struct {
	amount,price,fee *big.Float

}
func getKey(unit string,d int64) string {
return fmt.Sprintf("%d%s",d,unit)
}
//c:=big.NewFloat(0).Sub(big.NewFloat(Tod.LBalance),big.NewFloat(Pricebooks.SPB.Book[p] * p))
//c.SetPrec(c.MinPrec())
//Tod.LBalance,_ = c.Float64()
////Tod.LBalance -= Pricebooks.SPB.Book[p] * p
//c=big.NewFloat(0).Sub(big.NewFloat(Tod.Amount),big.NewFloat(Pricebooks.SPB.Book[p]))
//c.SetPrec(c.MinPrec())
//Tod.Amount,_=c.Float64()
////Tod.Amount -= Pricebooks.SPB.Book[p]

const ( Fee=0.000141)
func main() {
	//a:=Order{
	//	price:  0.5,
	//	amount: 0.008,
	//}
	//b:=O{
	//	price:  big.NewFloat(0.5),
	//	amount: big.NewFloat(0.008),
	//}
	//a.amount-=Fee*0.0001
	//b.amount.Sub(b.amount,big.NewFloat(Fee).Mul(big.NewFloat(Fee),big.NewFloat(0.0001)))
	//fmt.Println(a)
	//fmt.Println(b.amount.MinPrec())
	//b.amount=b.amount.SetPrec(b.amount.MinPrec())
	//fmt.Println(b.amount,b.price,b.amount.Prec())

	//fmt.Println(a.amount)
	//c:=big.NewFloat(0).Sub(big.NewFloat(a.amount),big.NewFloat(1)).SetMode(big.ToNearestAway)
	//fmt.Println(c, c.Prec())
	//c.SetPrec(c.Prec())
	//la,_:=c.Float64()
	//c.SetPrec(c.MinPrec())
	//fmt.Println(c.SetFloat64(la),c.MinPrec())
	//a.amount,_=c.Float64()
	//fmt.Println(a.amount)

	vi:=big.NewFloat(0).Sub(big.NewFloat(0.005),big.NewFloat(0.000000000015))
	vi.SetPrec(vi.MinPrec())
	fmt.Println(vi," vi value")
	//c:=big.NewFloat(0).Sub(big.NewFloat(a.amount),big.NewFloat(Fee))
	//fmt.Println("c: ",c)
	//c=big.NewFloat(0).Add(big.NewFloat(0.005),big.NewFloat(0.000000000015))
	//c.SetPrec(c.MinPrec())
	//x,damn:=c.Float64()
	//fmt.Println(c,"c: ",x,damn)
	//d:=big.NewFloat(0).Mul(big.NewFloat(a.amount),big.NewFloat(Fee))
	//d.SetPrec(d.MinPrec())
	//fmt.Println(d,c)
	//suis:=time.Now()
	//str:=getKey("sec",10)
	//fmt.Println("t: ",time.Since(suis)," s: ",str)
	//suis=time.Now()
	//fmt.Println("t: ",time.Since(suis)," s: ",str,str==str)
}
