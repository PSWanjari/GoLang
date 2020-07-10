package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	fmt.Println("I have selected the number let's see if you can guess it!")
	targetNumber := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)
	var f int = 1
	for x := 0; x < 10; x++ {
		fmt.Println(10-x,"Chance left!Guess a number:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if number > targetNumber {
			fmt.Println("opps! Your guess is High")
		}
		if number < targetNumber {
			fmt.Println("opps! Your guess is Low")
		}
		if number == targetNumber{
		    f=0
		    fmt.Println("Hurray !! you guess the correct number ",targetNumber )
		    break
		}
	}
	if f==1{
	    fmt.Println("Sorry You lose the game")
	}

}
