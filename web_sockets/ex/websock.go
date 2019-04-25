package ex

import (
	"net/http"
	"github.com/gorilla/websocket"
	"log"
	"github.com/julienschmidt/httprouter"
)


type Entries struct {
	Price ,Amount,Total float64
}

var upgrader =websocket.Upgrader{
	WriteBufferSize:1024,
	ReadBufferSize:1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func OBUpdater(w http.ResponseWriter,r *http.Request,_ httprouter.Params) {

	conn,err:=upgrader.Upgrade(w,r,nil)
	if err!=nil{
		log.Println("Error in connection")
		return
	}
	defer conn.Close()
	go Reader(conn)
	notifier(conn)

}

func Reader(con *websocket.Conn)  {
	for{
		_,_,err:=con.ReadMessage()
		if err!=nil{
			log.Println("Write: ",err)
			break
		}
	}
}

func notifier(con *websocket.Conn)  {
	for{
		i:=0.0
		v:=Entries{
			Price:25+i,
			Amount:5,
			Total:25*5,
		}
		err:=con.WriteJSON(&v)
		if err!=nil{
			log.Println("Write: ",err)
		break
		}
	}

}


