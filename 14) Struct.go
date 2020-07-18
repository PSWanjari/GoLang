package main

import (
	"fmt"
)

//program to store emplpoyee deatils
//structure is used inorder to store the value in the struct
func main() {
	fmt.Println("Pratik Wanjari")
	//creating struct to hold the value of the struct
	var employee struct {
		name       string
		salary     float64
		experiance int
		department string
	}
	employee.name = "Pratik"
	employee.salary = 13030300.292
	employee.experiance = 2
	employee.department = "IT"
	fmt.Println(employee)
	fmt.Printf("%0.3f ", employee.salary)

	
}
