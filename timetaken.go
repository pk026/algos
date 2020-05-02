package main
import "os"
import "fmt"
import "math"
import "bufio"
import "strconv"
import "strings"

func timeTaken(word_dict map[int][2]int, input string) int {
    var timeTakenToType float64
    var current_row, current_col float64
    var first rune
    for _, c := range input {
        first = c
        break
    }
    if rowcol, ok := word_dict[int(first)]; ok {
        current_row, current_col = float64(rowcol[0]), float64(rowcol[1])
    } else {
        return -1
    }
    for _, char := range input[1:] {
        if rowcol, ok := word_dict[int(char)]; ok {
            row, col := float64(rowcol[0]), float64(rowcol[1])
            timeTakenToType += math.Abs(current_row - row) + math.Abs(current_col - col)
        } else {return -1}
    }
    return int(timeTakenToType)
}
func main() {
    word_dict := make(map[int][2]int)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    row_col := scanner.Text()
    var n  int
    for index, value := range strings.Split(row_col, " ") {
        int_value, _ := strconv.Atoi(value)
        if index == 0 {
            n = int_value
        }
    }
    for i := 1; i <= n; i++ {
        scanner.Scan()
        for pos, char := range scanner.Text() {
            rowcol := [2]int{i, pos}
            word_dict[int(char)] = rowcol
        }
    }
    scanner.Scan()

    fmt.Printf("%v\n", timeTaken(word_dict, scanner.Text()))
}