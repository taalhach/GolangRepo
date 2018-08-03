package main
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)
var db =make(map[int][]int)
var i=0

type MatchedOrders struct {
	price int
	matched []int
}


func Router() *httprouter.Router{
	router:=httprouter.New()
	router.GET("/Match",Handler)
	return router
}

func Handler(w http.ResponseWriter,req *http.Request,_ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte(MatchOrders()))


}

func MatchOrders()string {
	ch:=make(chan MatchedOrders)
	a:=[]int{1,2,3,4,5,6}
	go Executor(ch,MatchedOrders{price:i,matched:a})
	go DB(ch)
	i++

	return "Accepted"
}

func Executor(ch  chan MatchedOrders,c MatchedOrders) {
	log.Println(c," connected")
	ch<-c

}

func DB(c chan MatchedOrders)  {

	x:=<-c
	db[x.price]=x.matched
	log.Println(db)

}

func main() {
	http.ListenAndServe(":8080",Router())
}