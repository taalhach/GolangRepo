package main

import (
	"github.com/taalhach/01_GolangRepo/02_Package_Example/Import_packagge"
	"fmt"
)
type Customer Import_packagge.Customer

func main() {

	customer1:=Import_packagge.Customer{"talha",5}
	customer1.GetCustomerDetails()
	customer2:=Customer{"Ali",25}
	customer2.CustomerInfo()




	//fmt.Println("calling variables directly due to package level scope")
	//fmt.Print("name: "+Import_packagge.Name+"\n username: " + Import_packagge.Username+"\n")
	//fmt.Println("typical getter like other programming languages")
	//fmt.Println("password: " +Import_packagge.GetPass())
	//fmt.Println("getting password from file")
	//fmt.Println("password: " +Import_packagge.GetPassFromFile())
	//

	}
func (cus *Customer)CustomerInfo()  {
	fmt.Println("Customer Name: ",cus.Name,"Customer id: ",cus.Id)

}