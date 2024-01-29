package main

import "fmt"

type Person struct {
    name string
    age  int
}

func newPerson(name string) *Person {
    p := Person{name: name}
    p.age = 42
    return &p
}

func main() {
    fmt.Println(Person{"Bob", 20})

    person1 := newPerson("Jon");
    fmt.Println(person1)

    person1.age = 30
    fmt.Println(person1)
}
