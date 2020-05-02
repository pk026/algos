package main

import "os"
import "fmt"
import "bufio"
import "strconv"
import "strings"

func equalSumPossible(int_list []int, sum, index int) bool {
    if index >= len(int_list) -1 {
        return false
    }
    if int_list[index] == sum {
        return true
    }
    return (equalSumPossible(int_list, sum - int_list[index], index + 1) || equalSumPossible(int_list, sum, index + 1))
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    var int_list []int
    var total_sum int
    for _, value := range strings.Split(scanner.Text(), "") {
        int_value, _ := strconv.Atoi(value)
        int_list = append(int_list, int_value)
        total_sum += int_value
    }
    if total_sum % 2 != 0 {
        fmt.Printf("%v\n", false)
    } else {
        fmt.Printf("%v\n", equalSumPossible(int_list, total_sum/2, 0))
    }
}