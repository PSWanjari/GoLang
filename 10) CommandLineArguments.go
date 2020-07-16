package main

import (
	"fmt"
	"os"
	"strconv"
)

//Program that take input from the command line arguments and print the Average of the number

func main() {
	numbers := os.Args[1:]
	average := 0
	for _, element := range numbers {

		n, _ := strconv.Atoi(element)
		average = average + n

	}
	fmt.Println(average / (len(numbers) - 1))

}
