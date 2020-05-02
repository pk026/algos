package main
import "fmt"
import "strconv"

func main() {
    fmt.Printf("Please enter a floating point number: ")
    var UserInput float64
    _, err := fmt.Scan(&UserInput)
    if err != nil {
    	fmt.Println(err)
    }
    fmt.Printf("You have entered : %f \n", UserInput)
    trancated := int(UserInput)
    fmt.Println("Alternative output ", strconv.FormatFloat(UserInput, 'f', 6, 64))
    fmt.Printf("%d", trancated)
}