package main

import (
	"github.com/gorilla/websocket"
	"log"
	"fmt"
	"net/url"
	"flag"
)
var addr = flag.String("addr", "localhost:3000", "http service address")

func main() {
	u:= url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	c,_,err:=websocket.DefaultDialer.Dial(u.String(),nil)
	if err!=nil{
		log.Println("fatal ",err)
		return
	}
	go func(con *websocket.Conn) {
		err:=c.WriteMessage(1,[]byte("Hey there"))
		if err!=nil {
			log.Println("write ",err)
		}
		fmt.Println("done writing")

	}(c)
	i:=0
	defer c.Close()
	for i<5 {
		_,msg,err:=c.ReadMessage()
		if err!=nil{
			log.Println("reading ",err)
			return
		}
		fmt.Println(string(msg))
		i++
	}
}


	//error indicates that the connection was closed after a partial message was read.
	//Does the Android client API allow the application to close the connection while a
	//message is in flight?