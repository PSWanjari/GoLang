package main
import("fmt"
        "log"
        )
//In go lang return type are writen at the end in parenthesis if manay 
//creating a simple calculator for cost of two pens
func CalculateCost(costOfPen1 float64,costOfPen2 float64) (float64,error){
    //checking if the parameters are valid
    if costOfPen1<0{
        return 0.0, fmt.Errorf("The costOfPen1 i.e %0.2f is Invalid",costOfPen1)
    }
    if costOfPen2<0{
        return 0.0 , fmt.Errorf("The costOfPen2 ie 0.2f is Invalid",costOfPen2)
    }
    return costOfPen1+costOfPen2,nil
} 

func main(){
    cost,err := CalculateCost(12.59876,-2.5)
    if err!=nil{
        fmt.Println(err.Error())
        log.Fatal(err)
    }
    fmt.Printf("The cost of pen is %0.2f",cost)
}
