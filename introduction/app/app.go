package main

import "fmt"

// declaring the global varibales
var greetings = "Hello,"

func main() {
    // fmt.Println("hello world")
    customer := getData(2)
    fmt.Println(greetings + customer)
}

// func getData(variable dataType_var)(outputs){}
func getData(customerId int)(customer string){

    if customerId == 1{
        // traditional way of initializing 
    var firstName = "Praveen"
    lastName := "Hosamani" //:= initializing purpose
    // var fullName string

    return firstName + " " + lastName
    } else if customerId == 2 {
        return "Sharada Hosamani"
    } else {
        return " No customer name!!"
    }
    
}