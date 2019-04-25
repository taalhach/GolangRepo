package main




type Queue struct {
	MKT_BUY 	chan engineMOrder
	MKT_Sell 	chan engineMOrder
	LMT_BUY 	chan engineLOrder
	LMT_Sell 	chan engineLOrder
}

type PCHANS map[string]*Queue

func (p PCHANS)RegiterPair(pair string)  {

	if p[pair]!=nil{
		return
	}
	p[pair]=&Queue{
		MKT_BUY:make(chan engineMOrder),
		MKT_Sell:make(chan engineMOrder),
		LMT_BUY:make(chan engineLOrder),
		LMT_Sell:make(chan engineLOrder),
	}

}


