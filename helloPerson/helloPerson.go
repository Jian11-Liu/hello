package main

import "fmt"

const helloPrefix = "Hello, "

func HelloPerson(person string) string {
    return helloPrefix + person
}

func main() {
    fmt.Println(HelloPerson("LiuJian"))
}
