/*
Data.txt file conatin

Pratik
Rahul
Shweta
Shweta
Pratik
Vaish
Vaish
Prajakta
Pratik
Pratik
Rahul
Rahul
*/
//main code starts
package main

import (
	"fmt"
	"log"
	"readfile"
)

func main() {
	content, err := readfile.ReadFile("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(content)
	//map
	var mymap map[int]int
	fmt.Println(mymap)
	mymap = make(map[int]int)

	mymap[1] = 100
	mymap[2] = 200
	fmt.Println(mymap[3])

	//Secound return value from a map is boolean whcih is optional shows that if the value is present or not in the map
	fmt.Println("----------Check if key is present-------------")
	val, isok := mymap[3]
	fmt.Println(isok) //returns false
	fmt.Println(val)  // returns o

	//if you only whants to check if a particular key is present in the map you can just omit 1 value
	//check if key 4 is present in map
	_, isPresentInMap := mymap[4]
	if isPresentInMap {
		fmt.Println("The key is present")
	} else {
		fmt.Println("The key is not present")
	}

	//Deleting the element from the map
	//Just pass the delete function two things: the map you want to delete a key from, and the key you want deleted
	fmt.Println("-------Delete key from map-------------")
	fmt.Println(mymap)
	delete(mymap, 1)
	fmt.Println(mymap)

	fmt.Println("-------Updating key from map-------------")
	mymap[2] = 300
	fmt.Println(mymap)
	//We Have to read the file and calculate the votes
	//Declaring the map named voting_map
	//if key is not present in map the it return nil
	var winnerName string
	maxcount := 0
	voting_map := make(map[string]int)
	for _, item := range content {
		voting_map[item]++
	}
	fmt.Println(voting_map)

	//Itarating the map
	//We can use the for ... range loop to iterrate the map
	//syntax: for key,value := range map {....}
	fmt.Println("-----------------------")
	fmt.Println("|  Election result    |")
	fmt.Println("-----------------------")
	for key, value := range voting_map {
		fmt.Printf("|%10s : %d%7s|\n", key, value, "")
		if value > maxcount {
			maxcount = value
			winnerName = key
		}
	}
	fmt.Println("-----------------------")
	fmt.Printf("| Winner: %s      |\n", winnerName)
	fmt.Println("-----------------------")
}



