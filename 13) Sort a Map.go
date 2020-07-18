package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Sorting the map")
	mymap := make(map[string]int)
	mymap["vaish"] = 4
	mymap["pratik"] = 1
	mymap["rahul"] = 2
	mymap["shweta"] = 3

	fmt.Println(mymap)
	//While Ittarateing the map the values come in random
	//In order to sort the map first we need to get the key in slice and then sort the slice
	myslice := make([]string, 0)
	//itarate through the map and fill the slice with map key which we need to display in assc order
	for key, _ := range mymap {
		myslice = append(myslice, key)
	}
	fmt.Println(myslice)

	//sorting the slice using sort package
	sort.Strings(myslice)
	fmt.Println(myslice)
}
