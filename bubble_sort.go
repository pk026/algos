package main
import "fmt"
import "log"

func UserInput(input_slice *[]int) {
    var int_input int
    for input_count := 10; input_count > 0; input_count-- {
        fmt.Printf("Please enter upto: ")
        _, err := fmt.Scan(&int_input)
        if err != nil {
            log.Fatal(err)
        }
        *input_slice = append(*input_slice, int_input)
    }

}

func Swap(array []int, index_i, index_j int) {
    temp := array[index_i]
    array[index_i] = array[index_j]
    array[index_j] = temp
}

func BubbleSort(array []int) {
    if len(array) > 1 {
        for _, _ = range array {
            index_i := 0
            index_j := 1
            for index_j < len(array) {
                if array[index_i] > array[index_j] {
                    Swap(array, index_i, index_j)
                }
                index_i++
                index_j++
            }
        }
    }
}

func main() {
    integers := make([]int, 0, 0)
    UserInput(&integers)
    BubbleSort(integers)
    fmt.Printf("Sorted %v \n", integers)
}
