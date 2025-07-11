package main

import (
    "fmt"
)

func main() {
    // var numbers = []int{1, 1, 2, 3, 5, 8}
    // for i, num := range numbers {
	// 	fmt.Print(num)
	// 	if i < len(numbers)-1 {
	// 		fmt.Print(" ")
	// 	}
	// }
    numbers := []int{}
    var totalVariables int
    var eachNum int
    var totalSum int
    var operation int
    // var sum int
    fmt.Print("How many numbers do you want to work with?: ")
    fmt.Scanln(&totalVariables)
    for i := totalVariables; i >= 1; i-- {
        fmt.Print("Write a number: ")
        fmt.Scanln(&eachNum)
        numbers = append(numbers, eachNum)
        if i >= 2 {
            fmt.Println("What do you want to do?")
            fmt.Println("1. Addition")
            fmt.Println("2. Subtraction")
            fmt.Println("3. Multiplication")
            fmt.Println("4. Division")
            fmt.Scanln(&operation)
        }
    }
    for i := range numbers {
        switch operation {
        case 1:
            totalSum = totalSum + numbers[i]
        case 2:
            totalSum = totalSum - numbers[i]
        case 3:
            totalSum = totalSum * numbers[i]
        case 4:
            totalSum = totalSum / numbers[i]
        }
    }
    fmt.Print("The total is: ", totalSum, "\n")
}