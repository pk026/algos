package main
import "fmt"
import "sort"
import "os"
import "strconv"

func main() {
    int_list := make([]int, 0, 3)

    var UserInput string
    for {
        fmt.Printf("Please enter a int to add into int slice or 'x' to abort the program: ")
        _, err := fmt.Scan(&UserInput)
        if err != nil {
        	fmt.Println(err)
        }
        // if UserInput == "X" {
        //     os.Exit(3)
        // }
        var new_num int
        if new_num, err = strconv.Atoi(UserInput); err == nil {
            fmt.Printf("looks like a number.\n")
        } else if UserInput == "X" {
            os.Exit(3)
        }
        if new_num != 0 {
            int_list = append(int_list, new_num)
        }
        sort.Ints(int_list)
        fmt.Printf("UserInput %#v \n", int_list)
    }
}