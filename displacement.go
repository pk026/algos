package main
import "fmt"
import "log"

func userInput() (float64, float64, float64, float64){
    var acceleration, velocity, displacement, time float64;
    fmt.Printf("Please enter value for acceleration: ")
    _, err := fmt.Scan(&acceleration)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Please enter the value for initial velocity: ")
    _, err = fmt.Scan(&velocity)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Please enter the value for initial displacement: ")
    _, err = fmt.Scan(&displacement)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Please enter the value for time: ")
    _, err = fmt.Scan(&time)
    if err != nil {
        log.Fatal(err)
    }
    return acceleration, velocity, displacement, time
}

func GenDisplaceFn(acceleration, velocity, displacement float64) func(float64) float64 {
    cal_dis := func(time float64) float64{
        return 1.0/2 * acceleration * time * time + velocity * time + displacement
    }
    return cal_dis
}

func main() {
    acceleration, velocity, displacement, time := userInput()
    cald := GenDisplaceFn(acceleration, velocity, displacement)
    fmt.Println(cald(time))
}
