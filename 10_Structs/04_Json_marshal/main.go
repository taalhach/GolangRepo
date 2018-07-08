package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Movie struct {
	//fields name must be capitalised and exported by convention
	Name string
	Year int `json:"released"`
	Cast []string `json:"cast,omitempty"`
}

var moviesList =[]Movie{
	{Name:"Mission impossible ", Year:2000,Cast:[]string{"Tom cruise","Tom pegg"}},
	{Name:"Sicario",Year:2015,},
}

func main()  {
	mdata,err:=json.Marshal(moviesList)
	if err!=nil{
		log.Fatal("Error in marshalling")
	}
	fmt.Printf("%T\n","%v",mdata,"\n")
	fmt.Println(string(mdata))
	mdata,err=json.MarshalIndent(moviesList,""," ")
	fmt.Println(string(mdata))

}