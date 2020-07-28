package main

import "fmt"

/*
-------------Why we need type asseration---------------
 if a value of a concrete type is assigned to a variable with an interface type (including function parameters),
 then you can only call methods on it that are part of that interface,
regardless of what other methods the concrete type has

When you have a value of a concrete type assigned to a variable with an interface type,
a type assertion lets you get the concrete type back. It’s kind of like a type conversion.

syntacx:

var noise_maker NosieMaker = Robat("Iron man")
var robot Robat = noise_maker.(Robat)

In plain language, the type assertion above says something like “I know this
 variable uses the interface type NoiseMaker,
 but I’m pretty sure this NoiseMaker is actually a Robot.”
*/

//Making an interface noisemaker
type NoiseMaker interface {
	MakeNoise()
}

//type robat that satisfy the interface NoiseMaker
type Robot string

func (m Robot) RobatDance() {
	fmt.Println("Robat Dancing")
}

func (m Robot) MakeNoise() {
	fmt.Println("Hahahahaha Robot")
}

func main() {
	var noisemaker NoiseMaker
	noisemaker = Robot("Robot 1")
	noisemaker.MakeNoise()
	// but we cannot access the Robot walk method
	//so we need to do type casting which in go called type assertion
	robot := noisemaker.(Robot)
	robot.RobatDance()
}
