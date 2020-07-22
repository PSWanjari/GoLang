package main

import (
	"fmt"
	"log"
	"userdefinetype"
)

func main() {
	var todaysdate userdefinetype.Date
	err := todaysdate.SetMonth(0)
	if err != nil {
		log.Fatal(err)
	}
	err = todaysdate.SetDay(2)
	if err != nil {
		log.Fatal(err)
	}
	err = todaysdate.SetYear(23)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todaysdate)
}
