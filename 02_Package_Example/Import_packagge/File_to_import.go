package Import_packagge

import "fmt"

var Name string= "Resilient"
var Username string="res_789"
var pass string="pass123"

type Customer struct {
	Name string
	Id int
}
func (cus *Customer)GetCustomerDetails()  {
	fmt.Println("Customer Name: ",cus.Name," Customer ID: ",cus.Id)
}


