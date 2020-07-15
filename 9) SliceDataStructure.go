package main

import "fmt"

func main() {
	fmt.Println("Hello")
	myarray := [6]string{"A", "B", "C", "D", "E", "F"}
	slice_ds := make([]string, 5)
	slice_ds[2] = "bye"
	fmt.Printf("%s", myarray[0])
	for _, item := range slice_ds {
		fmt.Printf("%s", item)
	}
	//creating slice from the existing array
	//if we omit the start index the value of 0 will be used
	new_slice := myarray[1:3]
	for _, item := range new_slice {
		fmt.Println(item)
	}
	//slice is built on top of the underlaying array
	//it is nothing but a view of the array. when we call the make function the slice create an array and slice
	//is just the view of the array
	//**** any change in the slice will reflect in the change of the underlaying array and vice -a-versa***
	// better to create slices using make or a slice literal, rather than creating an array and using a slice operator on it

	array1 := [5]int{1, 1, 1, 1, 1}
	slice_of_array1 := array1[1:5]
	fmt.Println("Elements in array")
	for _, item := range array1 {
		fmt.Printf("%d ", item)
	}
	fmt.Println("\nElements in slice")
	for _, item := range slice_of_array1 {
		fmt.Printf("%d ", item)
	}
	slice_of_array1[3] = 4
	fmt.Println("\nElements in array")
	for _, item := range array1 {
		fmt.Printf("%d ", item)
	}
	fmt.Println("\nElements in slice")
	for _, item := range slice_of_array1 {
		fmt.Printf("%d ", item)
	}
	array1[2] = 10
	fmt.Println("\nElements in array")
	for _, item := range array1 {
		fmt.Printf("%d ", item)
	}
	fmt.Println("\nElements in slice")
	for _, item := range slice_of_array1 {
		fmt.Printf("%d ", item)
	}

	//add element in an slice
	//if the slice size became smaller all the elements will be copied to a new array
	slice_of_array1 = append(slice_of_array1, 100)
	fmt.Println("\nElements in slice")
	for _, item := range slice_of_array1 {
		fmt.Printf("%d ", item)
	}
	//note : don’t have to worry about whether you have an empty slice or a nil slice.
	//You can treat them both the same, and your code will “just work”!

}
