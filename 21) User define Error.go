package main

//The error interface is show bellow
/*type error interface {
	Error() string
}
we need to create a error type that satisfy the error interface
*/
import (
	"fmt"
	"log"
	"strconv"
)

//creating the type OutOfStock error
type OutOfStockError struct {
	desc     string
	quantity int
}

func (m OutOfStockError) Error() string {
	return fmt.Sprintf("The item is out of stock. Desc: %s", m.desc)
}

func main() {
	//I want to buy 15 item
	item := 15
	ok, err := buy(item)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%d item buyed succesfully", ok)
	}
}

func buy(quantity int) (int, error) {
	if quantity > 10 {
		var outofstock OutOfStockError
		outofstock.desc = "you asked for " + strconv.Itoa(quantity) + " item"
		outofstock.quantity = quantity
		return -1, outofstock
	} else {
		return quantity, nil
	}
}
