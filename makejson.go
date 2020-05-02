package main
import "fmt"
import "encoding/json"

func main() {
    var name, address string
    fmt.Printf("Please enter the name: ")
    _, _ = fmt.Scan(&name)
    fmt.Printf("Please enter address: ")
    _, _ = fmt.Scan(&address)
    // fmt.Printf("Name: %s, Address: %s \n", name, address)
    person := make(map[string]string)
    person["name"] = name
    person["address"] = address
    pd, err := json.Marshal(person)
    if err == nil {
        fmt.Printf(string(pd))
    } else {
        fmt.Printf("error %v", err)
    }
}