package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main()  {
	fmt.Print("Enter percentage: ")
	reader:=bufio.NewReader(os.Stdin)
	per,_,_:=reader.ReadLine()
	percent:=string(per)
	percentage,err:=strconv.ParseFloat(percent,64)
	if err!=nil{
	fmt.Println("Enter valid percentage ",err)
		panic("")
	}
	if percentage<50{
		fmt.Println("Grade: F")

	} else if percentage>50 && percentage<=60{
	fmt.Println("Grade: D")

	}else if percentage>60 && percentage<=70{
	fmt.Println("Grade: C")

	}else if percentage>70 && percentage<=80 {
		fmt.Println("Grade: B")
	}else if percentage>80 && percentage<=90 {
		fmt.Println("Grade: A")
	}else if percentage>90 {
		fmt.Println("Grade: A+")
	}
}
