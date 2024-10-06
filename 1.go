package main

import "fmt"

func main() {
    var operationInput string
    toEncrypt := false
	toROT13 := false
    for {
        fmt.Println("\nSelect operation (1/2):")
        fmt.Println("1. Encrypt.")
        fmt.Println("2. Decrypt.")
        fmt.Scan(&operationInput)

        if operationInput == "1" {
            toEncrypt = true
            break
        } else if operationInput == "2" {
            toEncrypt = false
            break
        } else {
            fmt.Println("Invalid input. Please select 1 or 2")
        }
    }
    fmt.Println(toEncrypt)
	for {
        fmt.Println("\nSelect cypher (1/2):")
        fmt.Println("1. ROT13.")
        fmt.Println("2. Reverse.")
        fmt.Scan(&cypherInput)

        if cypherInput == "1" {
            toROT13 = true
            break
        } else if cypherInput == "2" {
            toReverse = false
            break
        } else {
            fmt.Println("Invalid input. Please select 1 or 2")
        }
    }
    fmt.Println(toEncrypt)
	fmt.Println(toROT13)
}