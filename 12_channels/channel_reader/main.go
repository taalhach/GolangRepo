package main

import (
	"bufio"
	"os"
	"fmt"
	"github.com/taalhach/01_GolangRepo/12_channels/channel_reader/dmain"
)

func main()  {
	go get()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter")
	t,_:=reader.ReadString('\n')
	fmt.Print("Read: ",t)

}
func get()  {
	for i:=0;i<10 ;i++  {
		dmain.Ch<-&dmain.D{
			Name:"talha",
		}
	}
}