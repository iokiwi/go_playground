package main

import "fmt"

func greet(name string) {
    fmt.Printf(fmt.Sprintf("Hello, %s!\n", name))
}

func main() {
    greet("simon")    
}
