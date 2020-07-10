package main
import(
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the student name:")
    input,_ := reader.ReadString('\n')
    name := strings.TrimSpace(input) 
    fmt.Println("name is ",name)
    fmt.Println("Enter the student roll no:")
    input,_ =reader.ReadString('\n')
    input = strings.TrimSpace(input)
    rollno,_ := strconv.Atoi(input)
    fmt.Println("Roll number is ",rollno)
    fmt.Println("Enter the grade ")
    input,_ = reader.ReadString('\n');
    input = strings.TrimSpace(input)
    grade,_ := strconv.ParseFloat(input,64)
    fmt.Println("grade is ",grade)
    if grade>=60{
        fmt.Println(name,"result is Pass")
    }else{
        fmt.Println(name ,"result is Failed")
    }
    
    
}
