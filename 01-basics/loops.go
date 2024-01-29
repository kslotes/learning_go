package main

import "fmt"

func main() {
    i := 1
    for i <= 3 {
        fmt.Println(i)
        i++
    }

    for k := 4; k <= 10; k++ {
        fmt.Println(k)
    }

    for {
        fmt.Println("loop")
        break;
    }

    for n := 1; n <= 10; n++ {
        if n % 2 != 0 {
            continue
        }
        fmt.Println(n)
    }
}
