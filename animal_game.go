package main

import "os"
import "fmt"
import "bufio"
import "strings"

type Animal interface {
    Eat()
    Move()
    Speak()
}

type Cow struct {
    eat     string
    move    string
    speak   string
}

func (c Cow) Eat() {
    fmt.Printf("%s \n", c.eat)
}

func (c Cow) Move() {
    fmt.Printf("%s \n", c.move)
}

func (c Cow) Speak() {
    fmt.Printf("%s \n", c.speak)
}

type Bird struct {
    eat     string
    move    string
    speak   string
}

func (c Bird) Eat() {
    fmt.Printf("%s \n", c.eat)
}

func (c Bird) Move() {
    fmt.Printf("%s \n", c.move)
}

func (c Bird) Speak() {
    fmt.Printf("%s \n", c.speak)
}

type Snake struct {
    eat     string
    move    string
    speak   string
}

func (c Snake) Eat() {
    fmt.Printf("%s \n", c.eat)
}

func (c Snake) Move() {
    fmt.Printf("%s \n", c.move)
}

func (c Snake) Speak() {
    fmt.Printf("%s \n", c.speak)
}

func add_newanimal(my_animals map[string]Animal, name string, animal_type string) {
    switch animal_type {
        case "cow":
            my_animals[name] = Cow{eat: "grass", move: "walk", speak: "moo"}
        case "bird":
            my_animals[name] = Bird{eat: "worms", move: "fly", speak: "peep"}
        case "snake":
            my_animals[name] = Snake{eat: "mice", move: "slither", speak: "hsss"}
    }
    fmt.Printf("Created it! \n")
}

func get_query(my_animals map[string]Animal, name string, act string) {
    animal, exits := my_animals[name]
    if exits == true {
        switch act {
            case "eat":
                animal.Eat()
            case "move":
                animal.Move()
            case "speak":
                animal.Speak()
        }
    }
}

func main() {
    my_animals := make(map[string]Animal)
    scanner := bufio.NewScanner(os.Stdin)
    for {
        user_input := ""
        command := ""
        name := ""
        act := ""
        fmt.Printf("> ")
        scanner.Scan()
        user_input = scanner.Text()
        if len(user_input) > 0 {
            if len(strings.Split(user_input, " ")) == 3 {
                for index, value := range strings.Split(user_input, " "){
                    if index == 0 {
                        command = value
                    }
                    if index == 1 {
                        name = value
                    }
                    if index == 2 {
                        act = value
                    }
                }
            } else {
                fmt.Printf("Invalid input")
                continue
            }
            if command == "newanimal" {
                add_newanimal(my_animals, name, act)
            }
            if command == "query" {
                get_query(my_animals, name, act)
            }
        }
    }
}
