package main
import "fmt"
import "os"
import "bufio"
import "strings"

type Name struct {
	fname string;
	lname string;
}

func main() {
	fmt.Printf("Please enter the name file: ")
	name_slice := make([]Name, 0, 0)
	var file_path string
	_, err := fmt.Scan(&file_path)
	file, err := os.Open(file_path)
	if err != nil {
		fmt.Printf("Exiting cz %v", err)
		os.Exit(3)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		name := strings.Split(line, " ")
		name_slice = append(name_slice, Name{fname: name[0], lname: name[1]})
	}
	for _, name := range name_slice {
		fmt.Printf("First name: %s Last Name: %s \n", name.fname, name.lname)
	}
}