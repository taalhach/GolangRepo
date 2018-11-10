package routers

import (
	"github.com/gorilla/websocket"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
	"time"
)
var allconnections= []*websocket.Conn{

}

type Response struct {
	Message string `json:"message"`

}

type Request struct {
	Name string `json:"name"`
	Password string`json:"password""`
}
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
	return true
	},
	ReadBufferSize:1024,
	WriteBufferSize:1024,

}

// handler to websocket request
func WebsocketConnection(w http.ResponseWriter,r *http.Request, params httprouter.Params)  {
	//http to websocket protocol upgrader
	con,err:=upgrader.Upgrade(w,r,nil)
	if err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request protocol shift error"))
		return
	}
	//adding user connection in array so that updates are sent to user
	allconnections=append(allconnections,con)
	// infinite reader that reads from user and logs message
	go func(con *websocket.Conn ) {
		for{
			// readJson reads json from user and save in v
			v:=Request{}
			err=con.ReadJSON(&v)
			if err!=nil{
				con.WriteMessage(1,[]byte("Error in reading "))
				con.Close()
				// deletes user from array so that no more updates are provided to that user
				deleteUser(con)
				return
			}
			log.Println(v,"fom client")
			// readMessage reads meesage of text form from user
			_,msg,err:=con.ReadMessage()
			if err!=nil{
				con.WriteMessage(1,[]byte("Error in reading "))
				return
			}
			log.Println(msg, "from routers")
		}
	}(con )
	// provider provides updates after every five seconds to all user's conneaction who are in array
	provider()

}

//deleteuser iterates array and deletes connection from array
func deleteUser(c *websocket.Conn)  {
	for i:=0;i<len(allconnections) ;i++  {
	if c==allconnections[i]{
		allconnections=append(allconnections[:i],allconnections[i+1:]...)
	}
	}

}
//NOTE: writeJson ReadJson WriteMessage WriteJson one type of reader and writer are used  
// writer to connection
func provider() {
	// golang channel that generates an event after every 5 seconds
	ch := time.Tick(5 * time.Second)
	// infinite writer
	for {
		for range ch {
			// range all users connections and provide updates
			for _, con := range allconnections {
				c := Response{
					Message:"Hey there",
				}
				// writes json to user
				err:=con.WriteJSON(&c)
				if err!=nil{
					con.WriteMessage(1,[]byte("Error in writing "))
					con.Close()
					deleteUser(con)
					return
				}
				// also can use write mesage
				err=con.WriteMessage(1,[]byte("Hey there"))
				if err!=nil{
					con.WriteMessage(1,[]byte("Error in writing "))
					con.Close()
					deleteUser(con)
					return
				}
			}

		}

	}

}

