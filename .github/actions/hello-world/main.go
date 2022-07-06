package main

import(
    "fmt"
    "os"
)

func main() {
    firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
    secondGreeting := os.Getenv("INPUT_SECONDGREETING")
    pathOfGreating := os.Getenv("INPUT_THIRDGREETING")
    
    fmt.Println("Hello First" + firstGreeting)
    fmt.Println("Path1 " + secondGreeting)
    fmt.Println("Path2 " + pathOfGreating)
    
    if secondGreeting != "" {
    fmt.Println("Hello " + secondGreeting)
    }
}
