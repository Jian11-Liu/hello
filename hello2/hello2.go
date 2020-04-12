package main

import "fmt"

const helloPrefix = "Hello, "

func Hello(person string) string {
    if person == "" {
        person = "World"
    }
    return helloPrefix + person
}

func main() {
    fmt.Println(Hello(""))
}
