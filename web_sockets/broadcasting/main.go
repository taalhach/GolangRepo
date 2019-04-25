package main

import (
	"encoding/json"
	"fmt"
	"github.com/block360/mengine"
	"github.com/block360/engine"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"
)

var allConnection []*websocket.Conn

type Trade struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Time   string  `json:"time"`
	Type   string  `json:"type"`
}

type entry struct {
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Total  float64 `json:"total"`
}
type OB struct {
	Asks         interface{} `json:"asks"`
	Bids         interface{} `json:"bids"`
	TradeHistory interface{} `json:"tradehistory"`
}

//var cha chan string
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WShandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//fmt.Println("called")
	//conn, err :=upgrader.Upgrade(w,r,nil)
	//if err!=nil{
	//	fmt.Println("Error in upgrading")
	//	return
	//}
	//allConnection=append(allConnection,conn)
	//  func(c *websocket.Conn) {
	//	for{
	//		defer c.Close()
	//		_,_,err:=c.ReadMessage()
	//		if err!=nil{
	//			for i,_:=range allConnection{
	//				if allConnection[i]==conn{
	//					allConnection=append(allConnection[:i],allConnection[i+1:]...)
	//					break
	//				}
	//			}
	//			break
	//		}
	//	}
	//	}(conn)
}

func init() {
	//cha=make(chan string)
	//damnEngine.TradeStream=make(chan damnEngine.TradeHistory)
	//go THbroadcasting()
	go func() {
		//for data:=range cha{
		//	for _,con:=range allConnection{
		//		fmt.Println(data,"c")
		//		con.WriteMessage(websocket.TextMessage,[]byte(data))
		//	}
		//}
	}()
}

func GetPeparedVals() ([]entry, []entry) {
	a := []entry{}
	b := []entry{}
	go func() {
		for p, am := range engine.SObEnttries.PriceAmounts {
			fmt.Println(p, am)
			a = append(a, entry{
				Price:  p,
				Amount: am,
				Total:  p * am,
			})
		}
	}()
	for p, am := range engine.BObEnttries.PriceAmounts {
		b = append(b, entry{
			Price:  p,
			Amount: am,
			Total:  p * am,
		})
	}

	return a, b
}
func Producer() {
	//fmt.Println("called p")
	//i:=0
	//ch:=time.Tick(1*time.Second)
	//for range ch{
	//	cha<-"Message: "+strconv.Itoa(i)
	//	fmt.Println(i)
	//	i++
	//}
}

func Orderbook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		return
	}
	allConnection = append(allConnection, con)

	go func(c *websocket.Conn) {
		defer c.Close()
		a := OB{}
		for {
			err := c.ReadJSON(&a)
			if err != nil {
				for i, _ := range allConnection {
					if allConnection[i] == c {
						allConnection = append(allConnection[:i], allConnection[i+1:]...)
						break
					}
				}
				break
			}
		}
	}(con)
	provider()
}
func getDummyTrades() interface{} {
	trades := []Trade{}
	i := 1
	for i < 20 {
		t := time.Now()
		trd := Trade{
			Price:  float64(100 + i),
			Amount: float64(1 + i),
			Time:   strconv.Itoa(t.Hour()) + ":" + strconv.Itoa(t.Minute()) + ":" + strconv.Itoa(t.Second()),
		}
		if i%2 == 0 {
			trd.Type = "buy"
		} else {
			trd.Type = "sell"
		}
		trades = append(trades, trd)
		i++
	}
	return trades
}
func OrderbookEntry(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	go func(con *websocket.Conn) {
		for {
			_, _, err = con.ReadMessage()
			if err != nil {
				log.Println("Disconnected")
				return
			}
		}
	}(con)
	if err := con.WriteJSON(&OB{
		Asks:         Mengine.SPB.GetOrderbook(),
		Bids:         Mengine.BPB.GetOrderbook(),
		TradeHistory: getDummyTrades(),
	}); err != nil {
		log.Println("Error in writing")
		return
	}
	{

	}
	orderbookvalues(con)
}

func orderbookvalues(con *websocket.Conn) {
	ch := time.Tick(10 * time.Second)
	for range ch {
		con.WriteMessage(1, []byte("Hey there we are still connected"))
	}
}

func provider() {

	ch := time.Tick(5 * time.Second)

	for {
		for range ch {
			for _, con := range allConnection {
				c := OB{}
				c.Asks, c.Bids = GetPeparedVals()
				w, err := con.NextWriter(websocket.TextMessage)
				if err != nil {
					fmt.Println("encoding error")
					return
				}
				err = json.NewEncoder(w).Encode(&c)
				if err != nil {
					fmt.Println("sucks")
					w.Close()
					return
				}
			}

		}

	}

}

//func THbroadcasting()  {
//	for th:=range damnEngine.TradeStream{
//		fmt.Println("Here in bd")
//		time.Sleep(1*time.Second)
//		for _,con:=range allConnection{
//			con.WriteJSON(th)
//		}
//	}
//
//}

func main() {
	//	go Producer()
	engine.LoadOrder()
	//go damnEngine.SellLimitOrder()
	//	damnEngine.BuyOrderBook.PrintOrderBook()
	//	damnEngine.SellOrderBook.PrintOrderBook()
	//fmt.Println(GetPeparedVals())
	router := httprouter.New()
	//fmt.Println("Buy: ",damnEngine.BObEnttries)
	//fmt.Println("Sell: ",damnEngine.SObEnttries)
	//fmt.Println("main")
	//router.GET("/ws",WShandler)
	//fmt.Println(timeStamp)
	//fmt.Println(time.Unix(time.Now().UnixNano() / int64(time.Millisecond),0).Clock())
	fmt.Println(time.Now().Hour(), ":", time.Now().Minute(), ":", time.Now().Second())
	router.GET("/ob", Orderbook)
	router.GET("/exchange", OrderbookEntry)
	http.ListenAndServe(":3000", router)
}
