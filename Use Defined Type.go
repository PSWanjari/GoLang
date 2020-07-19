package main

import (
	"fmt"
	"userdefinetype"
)

/****************************define a new type in go****************************

syntax ; type Mytype struct {
	--code here--
}
types are usually defined outside of any functions, at the package level.

If you’ve defined a type named car in the current package, and you declare a
variable that’s also named car,the variable name will shadow the type name, making it inaccessible.
*/
type employee struct {
	name   string
	salary float64
	dept   department
}

type department struct {
	deptId   int
	deptName string
}

func main() {

	fmt.Println("Type demo")

	var dept department
	var emp employee
	dept.deptId = 1
	dept.deptName = "IT"
	emp.name = "Pratik"
	emp.salary = 100
	emp.dept = dept
	doubleSalary(&emp)
	fmt.Println(emp)
	var worker1 userdefinetype.Worker
	// You can have a mixture of exported and unexported fields within a single struct type, if you want.
	worker1.Name = "Pratik"
	worker1.Id = 1010101
	worker1.Deptartment = "IT"
	//Accessing the inner structure from the outer structure
	//worker1.Adres.Country = "India"
	//worker1.Adres.City = "Pune"
	//worker1.Adres.State = "Maharashtra"
	//worker1.Adres.Street = "B-704 Shivraj nagar colony number 1"
	//worker1.Adres.Landmark = "Opp royal majesctic"
	//worker1.Adres.Pin = 411017

	//we can access the Adress struct i.e inner struct using the anonymus struct
	worker1.Country = "India"
	fmt.Println(worker1)
	//structy literals
	worker2 := userdefinetype.Worker{Name: "Rahul", Id: 128756, Deptartment: "Marketing", Password: "P@ssw0rd", Salary: 1000.232}
	fmt.Println(worker2)

}

//pass by reference of user define type
func doubleSalary(emp *employee) {
	emp.salary = emp.salary * 2.0
}
