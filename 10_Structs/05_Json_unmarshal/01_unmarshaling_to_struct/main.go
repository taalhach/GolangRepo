package main
import (
	"encoding/json"
	"log"
	"fmt"
)

var b = []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

type PersonalData struct {
	Name string
	Age int
	Parents []string
}
var pd PersonalData
func main()  {

	err:=json.Unmarshal(b,&pd)
	if err!=nil {
		log.Fatal("Error in Unmarshaling Personal data")
	}
	fmt.Println(string(b))
}