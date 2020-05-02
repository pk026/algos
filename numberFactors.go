package main
// Efficient program to print the number of factors of n numbers
import "fmt"
import "math"

func getFactorCount(number int) int {
	cnt := 0
	sqrt := int(math.Sqrt(float64(number)))
	for i := 1; i <= sqrt; i++ {
		if number % i == 0 {
			if number / i == i {
				cnt += 1
			} else {
				cnt += 2
			}
		}
	}
	return cnt
}

func numberFactors(input []int) []int {
	// memorize := make(map[int]int)
	var factorCountList []int
	for _, number := range input {
		factorCountList = append(factorCountList, getFactorCount(number))
	}
	return factorCountList
}

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1000000000000000000}
	out := numberFactors(input)
	fmt.Printf("%v \n", out)
}