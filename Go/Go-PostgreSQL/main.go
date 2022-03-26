package main

import "fmt"

type Person struct {
	firstName string
	age       int
}

func main() {
	var tim Person
	tim.firstName = "Tim"
	tim.age = 10
	fmt.Println(tim) // => {Tim 10}

	bob := Person{"Bob", 15}
	fmt.Println(bob) // => {Bob 15}

	sam := Person{age: 20, firstName: "Sam"}
	fmt.Println(sam) // => {Sam 20}

	mike := new(Person)
	mike.firstName = "Mike"
	mike.age = 25
	fmt.Println(mike) // => &{Mike 25}

	luke := &Person{"Luke", 30}
	fmt.Println(luke) // => &{Luke 30}
}
