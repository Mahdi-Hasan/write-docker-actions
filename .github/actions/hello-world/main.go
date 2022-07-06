package main

import(
    "fmt"
    "os"
)

func main() {
    firstGreeting := os.Getenv("INPUT_FIRSTGREETING")
    secondGreeting := os.Getenv("INPUT_SECONDGREETING")
    pathOfGreating := os.Getenv("INPUT_PathOfGreating")
    
    fmt.Println("Hello First" + firstGreeting)
    fmt.Println("Path " + pathOfGreating)
    
    if secondGreeting != "" {
    fmt.Println("Hello " + secondGreeting)
    }
}
