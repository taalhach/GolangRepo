package main

import (
	"fmt"
	"bufio"
	"os"
)

func main()  {
	fmt.Println("Enter grade: ")
	reader:=bufio.NewReader(os.Stdin)
	grade,_,_:=reader.ReadRune()
	userGrade:=string(grade)
	switch userGrade {
	case "F","f":
		fmt.Println("you obtained less than 50% marks in your exams")
	case "D","d":
		fmt.Println("you obtained less than 60% but greater than 50% marks")
	case "c","C":
		fmt.Println("you obtained less than 70% but greater than 60% marks")
	case "b","B":
		fmt.Println("you obtained less than 80% but greater than 70% marks")
	case "A","a":
		fmt.Println("you obtained 90% or greater than 80% marks")
	fallthrough
	case "":
		fmt.Println("Conratulation topper")
	default:
		fmt.Println("Enter valid grade A,B,C,D,F")



	}

}
