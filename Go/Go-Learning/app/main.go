package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	var mike Person
	mike.firstName = "Mike"
	mike.age = 20
	fmt.Println(mike) // => {Mike 20}

	bob := Person{"Tim", 30}
	fmt.Println(bob) // => {Tim 30}

	sam := Person{age: 15, firstName: "Sam"}
	fmt.Println(sam) // => {Sam 15}

	tim := new(Person)
	tim.firstName = "Bob"
	tim.age = 25
	fmt.Println(tim) // => &{Bob 25}
}
