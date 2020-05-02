package main
import "fmt"
// import "math"
//Minimum number of jumps to reach end
// Given an array of integers where each element represents the max number of steps
// that can be made forward from that element.
// Write a function to return the minimum number of jumps to reach the end of the array
// (starting from the first element). If an element is 0,
// then cannot move through that element.

func minJumps(array []int, l, h int) int {
	if h == l {
		return 0
	}
	min := 10000000000000
	for i := 1; i <= array[l]; i ++ {
		jumps := minJumps(array, i, h)
		if min < jumps + 1 {
			min = jumps + 1
		}
	}
	return min
}
func main() {
	array := []int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", minJumps(array, 0, len(array) - 1))
}

// package main

// import "os"
// import "fmt"
// import "bufio"
// import "strconv"
// import "strings"

// func main() {
//     scanner := bufio.NewScanner(os.Stdin)
//     scanner.Scan()
//     scanner.Scan()
//     input_list := scanner.Text()
//     var array []int
//     for _, value := range strings.Split(input_list, " ") {
//         integer, _ := strconv.Atoi(value)
//         array = append(array, integer)
//     }
//     for _, number := range array {
//         for i := 1; i <= number; i++ {
//             if i % 15 == 0 {
//                 fmt.Printf("FizzBuzz\n")
//             } else if i % 5 == 0 {
//                 fmt.Printf("Buzz\n")
//             } else if i % 3 == 0 {
//                 fmt.Printf("Fizz\n")
//             } else {
//                 fmt.Printf("%d\n", i)
//             }
//         }
//     }
// }