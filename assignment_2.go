package main

import (
    "fmt"
)

func main() {
    var input string
    fmt.Print("Enter a string: ")
    fmt.Scanln(&input)
    fmt.Println("You entered:", input)

	var str1, str2 string
    fmt.Print("Enter the first string: ")
    fmt.Scanln(&str1)
    fmt.Print("Enter the second string: ")
    fmt.Scanln(&str2)

    if len(str1) != len(str2) {
        fmt.Println("The strings are of different lengths.")
    } else {
        differences := make([]int, 0)
        for i := range str1 {
            if str1[i] != str2[i] {
                differences = append(differences, i)
            }
        }

        if len(differences) == 0 {
            fmt.Println("The strings are identical.")
        } else {
            fmt.Printf("The strings differ at the following positions: %v\n", differences)
        }
    }
}
