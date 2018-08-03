package main
import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
)
var db =make(map[int]string)
var i=0
func Router() *httprouter.Router{
	router:=httprouter.New()
	router.GET("/home",Handler)
	return router
}

func Handler(w http.ResponseWriter,req *http.Request,_ httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("Hey there"))
	ch:=make(chan string)
	go LogResponse(ch,req.RemoteAddr)
	go DB(ch)
}

func LogResponse(ch  chan string,c string) {
	log.Println(c," connected")
	ch<-c
}

func DB(c chan string)  {

	db[i]=<-c
	log.Println(db)
	i++
}

func main() {
	http.ListenAndServe(":8080",Router())
}