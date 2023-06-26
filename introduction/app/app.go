package main

import "fmt"

// declaring the global varibales
var greetings = "Hello,"

func main() {
    // fmt.Println("hello world")
    customers := getData()
    fmt.Println(customers)
    fmt.Println(len(customers))
}

// func getData(variable dataType_var)(outputs){}
func getData()(customers []string){
    // array
    customers = []string {"Praveen", "Sharada", "Pragati", "Prashant"}

    customers = append(customers, "Shruti")

    for x := 0; x < len(customers); x++ {
        fmt.Println(customers[x])
    }

    return customers
}