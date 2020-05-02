package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "strconv"

func iterativeFactorial(n int64) int64 {
	factorial := int64(1)
	for i := n; i > 0; i-- {
		factorial = factorial * i
	}
	return factorial
}

func recursiveFactorial(n, factorial int64, memorize map[int64]int64) int64 {
	if n <= 1 {
		return factorial
	}
	if val, ok := memorize[n]; ok {
		fmt.Printf("Saving cal for %d", n)
		return val
	}
	result := n * recursiveFactorial(n -1, factorial, memorize)
	memorize[n] = result
	return result
}

func main() {
	fmt.Printf("Put the number n for which you wish to calculate factorial for: ")
	newScanner := bufio.NewScanner(os.Stdin)
	newScanner.Scan()
	input := strings.Trim(newScanner.Text(), " ")
	n, _ := strconv.ParseInt(input, 10, 8)
	// factorial := iterativeFactorial(n)
	memorize := make(map[int64]int64)
	factorial := recursiveFactorial(n, int64(1), memorize)
	fmt.Printf("factorial for %d is %d\n", n, factorial)
}