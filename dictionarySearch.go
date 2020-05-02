package main

import "os"
import "fmt"
import "bufio"
import "strconv"
import "strings"

func main() {
	word_mapping := make(map[int][]map[int]bool)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, m int
	for index, value:= range strings.Split(scanner.Text(), " ") {
		int_value, _ := strconv.Atoi(value)
		if index == 0 {
			n = int_value
		}
		if index == 1 {
			m = int_value
		}
	}
	for i := 0; i< n; i++ {
		scanner.Scan()
		// word := scanner.Text()
		for index, char := range scanner.Text() {
			if count_slice, ok := word_mapping[int(char)]; ok {
				count_slice[index][i] = true 
			} else {
				word_dict := make(map[int]bool)
				word_dict[i] = true
				index_count := make([]map[int]bool, m, m)
				index_count[index] = word_dict
				word_mapping[int(char)] = index_count
			}
		}
	}
	temp_map := make(map[int]bool)
	matching_words := make(map[int]bool)
	for i := 0; i<m; i++ {
		matching_words[i] = true
	}
	scanner.Scan()
	for index, char := range scanner.Text() {
		if char != '?' {
			if mapping_list, ok := word_mapping[int(char)]; ok {
				for key, _ := range mapping_list[index] {
					if _, ok := matching_words[key]; ok {
						temp_map[key] = true
					}
				}
			}
			matching_words = temp_map
		}
	}
}
