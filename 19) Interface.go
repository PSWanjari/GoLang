package main

import (
	"fmt"
)

//interface with three method
type Myinterface interface {
	MethodWithoutParameter()
	MethodWithParamaeter(float64)
	MethodWithReturnValue() string
}

//interface with one method
type toggable interface {
	toggle()
}

//userdefine type switch
type Switch string

//method of type Switch to toggle switch
func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

// user define type wich have method from the 1st interface
type myType int

func (m myType) MethodWithoutParameter() {
	fmt.Println("Hello")
}
func (m myType) MethodWithParamaeter(value float64) {
	fmt.Println("Method2")
}

func (m myType) MethodWithReturnValue() string {
	fmt.Println("Method3")
	return "Method3"
}

func (m myType) NewFeatureAdded() {
	fmt.Println("I love this feature")
}
func main() {
	var value Myinterface
	value = myType(5)
	value.MethodWithoutParameter()
	var value2 myType
	value2.NewFeatureAdded()
	s := Switch("off")
	var toogle toggable = &s // pass as address for pointer
	toogle.toggle()
	toogle.toggle()
	toogle.toggle()
	toogle.toggle()
	toogle.toggle()
	toogle.toggle()
}
