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
func getData()(customer [4]string){
    // array
    customer[0] = "Praveen"
    customer[1] = "Sharada"
    customer[2] = "Pragati"
    customer[3] = "Prashant"

    return customer
}