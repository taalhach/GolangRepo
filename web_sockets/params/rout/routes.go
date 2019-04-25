package rout

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"time"
)

var upgrader =websocket.Upgrader{
	WriteBufferSize:1024,
	ReadBufferSize:1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ExchangeViewWs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	con, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request Error"))
	}
	c:=ps.ByName("name")
	fmt.Println(c)
	go writer(con,c)
	go reader(con)


}
func reader(conn *websocket.Conn)  {
	for{
		_,d,err:=conn.ReadMessage()
		if err!=nil{
			if websocket.IsUnexpectedCloseError(err,websocket.CloseGoingAway,websocket.CloseAbnormalClosure){
				log.Printf("error in reading %v",err)
			}
			conn.Close()
			break
		}
		log.Println("read, ",string(d))
	}
}

func writer(conn *websocket.Conn,name string) {
	for {
		time.Sleep(1*time.Second)
		err:=conn.WriteMessage(1,[]byte("hello "+name))
		if err!=nil{
			log.Printf("write error %v",err)
			return
		}
		fmt.Print("writing ")
	}
}
