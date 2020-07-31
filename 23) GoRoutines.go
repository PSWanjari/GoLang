package main

/*
In Go, concurrent tasks are called goroutines. Other programming languages
 have a similar concept called threads, but goroutines require less computer
 memory than threads, and less time to start up and stop, meaning you can run
 more goroutines at once.
 Syntax:

 go functioncall()

 go add("argument")

 we can’t use function return values in a go statement
 Go won’t let you use the return value from a function called with a go statement
 result = go add(a+b)//invalid
 But there is a way to communicate between goroutines: channels

*/
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	/*
		This func will call one after the other so it will take time
	*/

	responseSize("https://www.yahoo.com")
	responseSize("https://www.google.com")
	responseSize("https://www.facebook.com")
	responseSize("http://www.pratikwanjari.com")
	end := time.Now().Second
	println(int64(start))
}
func responseSize(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The size of " + url + " is:" + strconv.Itoa(len(body)) + " bytes")
}
