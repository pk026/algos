package main

import "os"
import "fmt"
import "strings"
import "bufio"

type Animal struct {
    eat     string;
    move    string;
    speak   string;
}

func (a *Animal) Eat() {
    fmt.Printf("%s \n", a.eat)
}

func (a *Animal) Move() {
    fmt.Printf("%s \n", a.move)
}

func (a *Animal) Speak() {
    fmt.Printf("%s \n", a.speak)
}

func BuildAnimalDir() map[string]Animal {
    animal_map := make(map[string]Animal)
    cow := Animal{eat: "grass", move: "walk", speak: "moo"}
    bird := Animal{eat: "worms", move: "fly", speak: "peep"}
    snake := Animal{eat: "mice", move: "slither", speak: "hsss"}
    animal_map["cow"] = cow
    animal_map["bird"] = bird
    animal_map["snake"] = snake
    return animal_map
}

func main() {
    var user_input string;
    var animal_name, activity string;
    animal_map := BuildAnimalDir()
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Printf("Please put animal and property> ")
        scanner.Scan()
        user_input = scanner.Text()
        // fmt.Printf("user_input %s", user_input)
        for index, value := range strings.Split(user_input, " "){
            if index == 0 {
                animal_name = value
            }
            if index == 1 {
                activity = value
            }
        }

        anim := animal_map[animal_name]
        if activity == "eat" {
            anim.Eat()
        }
        if activity == "move" {
            anim.Move()
        }
        if activity == "speak" {
            anim.Speak()
        }
    }
}