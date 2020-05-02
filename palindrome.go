package main

import "os"
import "fmt"
import "bufio"
import "strings"

func isPalindrome(start, end int, input string) bool {
	if (start >= end) {
		return true
	} else if (end - start == 1) {
		return input[start] == input[end]
	}
	if (input[start] != input[end]) {
		return false
	}
	return isPalindrome(start + 1, end -1, input)
}

func main() {
	fmt.Printf("Please enter string to check if it is palindrome: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Trim(scanner.Text(), " ")
	if isPalindrome(0, len(input) - 1, input) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}