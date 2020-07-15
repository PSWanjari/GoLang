package main

import "fmt"

func main() {

	fmt.Println("**********Define Array **************")
	//One way of defining array
	var myarray [5]string = [5]string{"Apple", "Banana", "Grapes", "Mango"}
	//another way of defining array if initaialize only we can use short descp
	alphabets := [3]string{"A", "B", "C"}
	fmt.Println(alphabets[0])
	fmt.Println(myarray[0])
	fmt.Println(myarray[2])
	fmt.Println(alphabets)
	fmt.Printf("%#v", myarray)

	//Iterate through the array by using a loop
	list := [5]string{"one", "two", "three", "four", "five"}
	for i := 0; i < len(list); i++ {
		fmt.Printf("%d is %s\n", i+1, list[i])
	}

	//Iterate Through the array using for...Range function
	for index, item := range list {
		fmt.Println(index, item)
	}

	for _, item := range list {
		fmt.Println(item)
	}

}
