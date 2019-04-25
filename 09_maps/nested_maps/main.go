package main

type balances map[string]float64
var Userbalances =make(map[string]balances)


func (b balances)init(){
	b["BTC"]=0.0
	b["ETH"]=0.0
	b["ERC"]=0.0
}


func main() {


}
