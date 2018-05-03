package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func main()  {

	//if you are using goland then file path given from root
	fileSrc,err:=os.Open("file.txt")
	if(err!=nil){
		fmt.Print(err)
		panic("404 file not found")

	}
	defer fileSrc.Close()
	fileData,err:=ioutil.ReadAll(fileSrc)
	if(err!=nil){
		panic("File reading error")
	}
	fmt.Printf("%T\n",fileData)
	fmt.Print(string(fileData))
}
