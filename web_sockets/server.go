package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/taalhach/01_GolangRepo/web_sockets/ex"
)

func main()  {
	router:=httprouter.New()
	router.GET("/orderbook",ex.OBUpdater)
	http.ListenAndServe(":3000",router)
}
