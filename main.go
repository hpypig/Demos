package main

import "fmt"

func main() {
    test(1, 1,1)
}

func test(a int, b ...int){
    fmt.Println(b)
}
