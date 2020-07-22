package main

import (
	"fmt"
	"log"
	"userdefinetype"
)

func main() {
	var todaysdate userdefinetype.Date
	//Setting the values ussing the setter mehod of the type
	err := todaysdate.SetMonth(2)
	if err != nil {
		log.Fatal(err)
	}
	err = todaysdate.SetDay(12)
	if err != nil {
		log.Fatal(err)
	}
	err = todaysdate.SetYear(1996)
	if err != nil {
		log.Fatal(err)
	}
	//getting the values using the getters
	fmt.Printf("Day : %d\nMonth : %d\nYear : %d", todaysdate.Day(), todaysdate.Month(), todaysdate.Year())
}
