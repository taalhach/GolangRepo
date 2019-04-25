package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"time"
)

var c chan string
func main()  {
	fmt.Println("hey")
	router :=httprouter.New()
	router.GET("/get",request)
	http.ListenAndServe(":9090",router)
}

func request(w http.ResponseWriter, r *http.Request, params httprouter.Params )  {
	w.Write([]byte("ok"))
	go caller()

}
func caller()  {
	go writer()
	go tester()
}
func tester()  {
	time.Sleep(1*time.Millisecond)
	fmt.Println("hey")
}

func writer()  {
	fmt.Println("called me")
	time.Sleep(1*time.Millisecond)
	c<-"hey there"

}

func init() {
	c=make(chan string)
	go func() {
		for s:=range c{
			time.Sleep(5*time.Millisecond)
			fmt.Println("i received: ",s)
		}
	}()
}