package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/taalhach/01_GolangRepo/web_sockets/simpleExample/routers"
	"net/http"
)

func main()  {
	router :=httprouter.New()
	router.GET("/notification", routers.WebsocketConnection)
	http.ListenAndServe(":9090",router)
}
