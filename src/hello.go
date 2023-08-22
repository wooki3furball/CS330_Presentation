package main

import(
        "fmt"
        "time"
)

func main() {
        for i := 0; ; i++ { // What does the inf loop print?
                if i%2 == 0 {
                foo1()
                fmt.Print("Even")
                time.Sleep(1 * time.Second)
                continue
                }

                fmt.Println("Hello World")
                defer foo1()
                foo2()
        }
}

func foo1() {
        fmt.Print("Gopher")
}

func foo2() {
        fmt.Print("Python")
}