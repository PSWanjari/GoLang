package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(mychannel chan int, b int, c int) {

	fmt.Println("inside add routine")
	time.Sleep(5 * time.Second)
	mychannel <- (b + c)
}
func mul(mychannel chan int, b int, c int) {

	fmt.Println("inside multiply routine")
	time.Sleep(30000)
	mychannel <- (b * c)
}

func Loadscreen(channel chan int) {
	//will load from 1 to 100
	for i := 1; i <= 100; i++ {
		//genrating a random number to genrate a scenario of loading
		//min 0 sec
		//max 3 sec
		responsetimefromserver := rand.Intn(2)
		time.Sleep(time.Duration(responsetimefromserver) * time.Second)
		channel <- i
	}
	channel <- -1

}
func main() {
	//creating a channel in order to pass the value
	mychannel := make(chan int)
	fmt.Println("In add function")
	go add(mychannel, 2, 3)
	fmt.Println("In mul function")
	go mul(mychannel, 2, 3)
	fmt.Println("waiting for add respsonse")
	myval := <-mychannel
	println(myval)
	fmt.Println("waiting for mul respsonse")
	myval = <-mychannel
	println(myval)
	//Loading program
	go Loadscreen(mychannel)
	for <-mychannel != -1 {
		fmt.Print("\u25AE")
	}

}
