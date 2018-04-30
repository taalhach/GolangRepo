package main;

import (
	"bufio"
	"os"
	"fmt"
)

func main()  {
	fmt.Print("Enter Name: ")
	reader:=bufio.NewReader(os.Stdin)
	str,_,_:=reader.ReadLine()
	alphabets:=getLength(string(str))
	fmt.Println(alphabets)

}
func getLength(str string) int {
return len(str)
}