package main

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"fmt"
	"time"
	"github.com/julienschmidt/httprouter"
)

type Entries struct {
	Price ,Amount float64
	Addr string
}
var upgrader =websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
func hadler( w http.ResponseWriter, r *http.Request,_ httprouter.Params){
	con,err:=upgrader.Upgrade(w,r,nil)
	if err!=nil {
		fmt.Println("Error in connection ")
		return
	}
	fmt.Println(con.RemoteAddr()," connected")
	defer con.Close()
	ch:=time.Tick(1*time.Second)
	z:=Entries{
		Price:100,
		Amount:0.02,
		Addr:con.RemoteAddr().String(),
	}

		for range ch{
			err=con.WriteJSON(&z)
			if err!=nil{
				log.Println("error in write",err)
				break
			}
		}

}

//var ws =new WebSocket("ws://localhost:3000/ws")
//ws.addEventListener("message",function(e){console.log(e)})
//ws.send("there he is ")
func main()  {
	router:=httprouter.New()
	router.GET("/ws",hadler)
	http.ListenAndServe(":3000",router)

}
