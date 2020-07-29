package main

/*
The defer keyword ensures a function call takes place even
if the calling function exits early, say, by using the return keyword.
Note: defer keyword itself can only be used with a function or method call.
*/
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func openFile(filename string) (*os.File, error) {
	fmt.Printf("opening the file %s\n", filename)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func closeFile(file *os.File) bool {
	fmt.Printf("closing the file %s\n", file.Name())
	if file != nil {
		return false
	}
	file.Close()
	return true
}

func getFolat64() ([]float64, error) {
	var numbers []float64
	file, err := openFile("data.txt")
	defer closeFile(file)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, number)

	}
	return numbers, nil

}
func main() {
	numbers, err := getFolat64()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(numbers)
}
