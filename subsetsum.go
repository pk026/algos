package main

import "fmt"
func moveZeroes(nums []int) {
	fmt.Printf("11111%v\n", nums)
    zero_count := 0
    non_zero_values := 0
    temp_fill := 0
    for _, value := range nums {
    	fmt.Printf("2222%v\n", nums)
        if value == 0 {
            zero_count += 1
            temp_fill += 1
        } else if temp_fill > 0 {
            nums[non_zero_values] = value
            non_zero_values += 1
            temp_fill -= 1
        }
    }
    length := len(nums)
    for zero_count > 0 {
        nums[length - zero_count] = 0
        zero_count -= 0
    }
    fmt.Printf("%v\n", nums)
}

func main() {
	// nums := []{}
	moveZeroes([]int{0, 1, 0, 3, 12})
}