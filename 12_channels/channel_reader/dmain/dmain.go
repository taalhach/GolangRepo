package dmain

import (
	"fmt"
	"time"
)

type D struct {
	Name string
}

func (d *D)Print()  {
	time.Sleep(1*time.Second)
	fmt.Println(d)
}
var Ch =make(chan *D)

func init() {
	go func() {
		for c:=range Ch{
			go c.Print()
		}
	}()
}