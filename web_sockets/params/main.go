package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/taalhach/01_GolangRepo/web_sockets/params/rout"
	"net/http"
)

func main() {
	router:=httprouter.New()
	router.GET("/data/:name",rout.ExchangeViewWs)
	http.ListenAndServe(":3030",router)
}
