package main

import "fmt"


func main() {
var num1 float64
var num2 float64
var operation string

fmt.Println("Insert your Calculation (format 1 + 1): ")
fmt.Scanf("%f%s%f", &num1,&operation,&num2)
var value float64
if operation == "+"{
value = num1 + num2
}else if operation == "-"{
   value=  num1 - num2
}else if operation == "*"{
   value =  num1 * num2
}else if operation == "/"{
    if num2 != 0{
        value = num1 / num2
    }else{
        fmt.Println("ERROR")
        return
    }
}else {
    fmt.Println("Please check the format example: 1 + 1 or use the correct operators -, /, +, *")
    return
}
fmt.Printf("Your calculation is: %.0f\n",value)
}
