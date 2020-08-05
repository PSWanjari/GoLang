package main

/*
The Go language supports first-class functions; that is, functions in Go are treated as “first-class citizens.”

In a programming language with first-class functions, functions can be assigned to variables, and then called from those variables.
*/
import (
	"fmt"
)

func sayHi() {
	fmt.Print("Hi from pratik!")
}

func takeFuncAndPrintTwice(Thfunc func()) {
	Thfunc()
	Thfunc()
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}
func mul(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}

func calculate(add func(int, int) int, sub func(int, int) int, mul func(int, int) int, div func(int, int) int) {
	a := 10
	b := 5
	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))

}

func main() {
	//asssigning a variable which can hold a function
	var myFunction func()
	myFunction = sayHi
	//Or using short Description
	myfunctiontemp := sayHi
	//calling the func variabble
	myFunction()
	myfunctiontemp()

	//*********************************************
	//passing the func as an argumrnt
	takeFuncAndPrintTwice(sayHi)

	//passing the func having argument
	calculate(add, subtract, mul, div)
}
