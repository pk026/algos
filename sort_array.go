package main

import "os"
import "fmt"
import "sync"
import "sort"
import "strconv"
import "bufio"
import "strings"

func sort_sub_array(wg *sync.WaitGroup, sub_array []int, sorted_list *[]int) []int {
    defer wg.Done()
    fmt.Printf("I am going to sort %v \n", sub_array)
    sort.Ints(sub_array)
    *sorted_list = sub_array
    return sub_array
}

func merge(array1, array2 []int, ) []int {
    var final_array []int
    array_length, i, j := len(array1) + len(array2), 0, 0
    for index := 0; index < array_length; index++ {
        switch true {
            case i == len(array1):
                final_array = append(final_array, array2[j])
                j += 1
            case j == len(array2):
                final_array = append(final_array, array1[i])
                i += 1
            case array1[i] < array2[j]:
                final_array = append(final_array, array1[i])
                i += 1
            case array1[i] >= array2[j]:
                final_array = append(final_array, array2[j])
                j += 1
        }
    }
    return final_array
}

func main() {
    fmt.Printf("Please import space separated list of integers to be sorted: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    user_input := scanner.Text()
    if len(user_input) > 0 {
        var list_input []int
        user_input = strings.TrimSpace(user_input)
        for _, value := range strings.Split(user_input, " "){
            if len(value) > 0 {
                value, err := strconv.Atoi(value)
                if err != nil {
                    fmt.Printf("Error: %v \n", err)
                }
                list_input = append(list_input, value)
            }
        }
        array_len := len(list_input)
        var wg sync.WaitGroup
        var sorted_sub_array1, sorted_sub_array2, sorted_sub_array3, sorted_sub_array4 []int
        if array_len >= 4 {
            sub_array_length := array_len / 4
            i := 0
            wg.Add(1)
            go sort_sub_array(&wg, list_input[i: i + sub_array_length], &sorted_sub_array1)
            i += sub_array_length
            wg.Add(1)
            go sort_sub_array(&wg, list_input[i: i + sub_array_length], &sorted_sub_array2)
            i += sub_array_length
            wg.Add(1)
            go sort_sub_array(&wg, list_input[i: i + sub_array_length], &sorted_sub_array3)
            i += sub_array_length
            wg.Add(1)
            go sort_sub_array(&wg, list_input[i: array_len], &sorted_sub_array4)
            wg.Wait()
        } else {
            for index, value := range list_input {
                if index == 0 {
                    sorted_sub_array1 = []int{value}
                }
                if index == 1 {
                    sorted_sub_array2 = []int{value}
                }
                if index == 2 {
                    sorted_sub_array3 = []int{value}
                }
                if index == 3 {
                    sorted_sub_array4 = []int{value}
                }
            }
        }
        // lets merge them
        merged_1_2 := merge(sorted_sub_array1, sorted_sub_array2)
        merged_3_4 := merge(sorted_sub_array3, sorted_sub_array4)
        final_array := merge(merged_1_2, merged_3_4)
        fmt.Printf("Sorted list %v \n", final_array)
    } else {
        fmt.Printf("Did you forget to put integers?")
    }
}
