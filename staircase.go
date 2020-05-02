package main
// Count ways to reach the n’th stair
import "os"
import "fmt"
import "bufio"
import "strconv"
import "strings"
 
func waysToClimb(n, m int64, memorise map[int64]int64) int64 {
	if n <= 2 {
		return n
	}
	if val, ok := memorise[n]; ok {
		return val
	}
	ways := int64(0)
	for i := int64(1); i <= m; i++ {
		ways += waysToClimb(n - i, m, memorise)
	}
	memorise[n] = ways
	return ways
}

func main() {
	fmt.Printf("Please enter number of staircase: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.Trim(scanner.Text(), " ")
	n, _ := strconv.ParseInt(input, 10, 8)
	fmt.Printf("Please enter number of staircase steps one could take: ")
	scanner.Scan()
	input = strings.Trim(scanner.Text(), " ")
	m, _ := strconv.ParseInt(input, 10, 8)
	memorise := make(map[int64]int64)
	fmt.Printf("Ways to climb %d \n", waysToClimb(n, m, memorise))
}