package main
import "fmt"
import "strings"

func main() {
    fmt.Printf("Please enter a String: ")
    var UserInput string
    _, err := fmt.Scan(&UserInput)
    if err != nil {
    	fmt.Println(err)
    }
    fmt.Printf("UserInput %s \n", UserInput)
    UserInput = strings.ToLower(UserInput)
    if (strings.HasPrefix(UserInput, "i") && strings.HasSuffix(UserInput, "n") && strings.Contains(UserInput, "a")) {
            fmt.Printf("%s \n", "Found!")
    } else {
        fmt.Printf("%s \n", "Not Found!")
    }
}