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

    fmt.Print("How many numbers do you want to work with?: ")
    fmt.Scanln(&totalVariables)

    fmt.Print("Write the first number: ")
	fmt.Scanln(&totalSum)
    
    for i := 2; i <= totalVariables; i++ {
        fmt.Println("What do you want to do?\n1. Addition\n2. Subtraction\n3. Multiplication\n4. Division")
        fmt.Scanln(&operation)

        fmt.Print("Write a number: ")
        fmt.Scanln(&eachNum)
        numbers = append(numbers, eachNum)

        switch operation {
        case 1:
            totalSum = totalSum + eachNum
        case 2:
            totalSum = totalSum - eachNum
        case 3:
            totalSum = totalSum * eachNum
        case 4:
            totalSum = totalSum / eachNum
        }
    }
    fmt.Print("The total is: ", totalSum, "\n")
}