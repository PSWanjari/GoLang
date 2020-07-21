package main

import (
	"fmt"
)

type Liters float64
type Gallon float64
type PratikNumber int

func main() {
	fmt.Println("Defined types")

	var carfule Liters
	var busfule Gallon
	mynumber := PratikNumber(3)
	mynumber.double()
	fmt.Println(mynumber)
	//short desc
	busfule = Gallon(4.5)
	carfule = Liters(4.5)
	fmt.Println(carfule, busfule)
	//we can only perform add mul div mul operation with same defined type even the underlaying type is same

	//method is a function which is associated with a values of a give type
	/*
		syntax:
		func (ReciverParameterName ReciverParameterType) SayHi(){
			---code here---
		}
	*/
	fmt.Println(busfule.ToLiter())
	fmt.Println(carfule.ToGallon())
}
func (value Liters) ToGallon() Liters {
	return value * Liters(0.4)
}
func (value Gallon) ToLiter() Gallon {
	return value * Gallon(2)
}

func (value *PratikNumber) double() {
	*value = *value * PratikNumber(2)
}
