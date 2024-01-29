package main

import "fmt"

func sum(nums ...int) int {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    var l = []int {1, 2, 3}

    fmt.Println(sum(1, 2))
    fmt.Println(sum(l...))
}

