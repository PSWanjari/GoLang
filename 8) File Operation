package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//get the file from the os
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	//add the file in the scanner
	scanner := bufio.NewScanner(file)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	average := 0
	index := 0
	//print the content of the file using scanner until the end of file
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		input := strings.TrimSpace(scanner.Text())
		number, _ := strconv.Atoi(input)
		average = average + number
		index++
	}
	println("Index ", index)
	println("Average", average)
	average = average / index

	//close the file and check for error
	err = file.Close()

	if err != nil {
		log.Fatal(err)
	}
}
