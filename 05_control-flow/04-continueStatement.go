package main
import (
	"bufio"
	"os"
	"fmt"
)
func main() {
	nameWithoutVowel:=""
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name")
	userName,_:=reader.ReadString('\n')
	userName=string(userName)
	vowel:=[]rune{'a','e','i','o','u'}

	for _,nameChar:=range userName{
		isVowel:=false
		for _,vow:= range vowel {
			if(vow==nameChar){
				isVowel=true
			}
		}
		if(isVowel){

			continue
		}
		nameWithoutVowel+=string(nameChar)
	}
	fmt.Println("Name without vowels: " ,nameWithoutVowel)
}
