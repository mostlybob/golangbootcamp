package main

import "fmt"

func main() {

	//2.1 Variables & inferred typing
	var (
		name, location     string
		age      int
	)

	name = "Prince Oberyn"
	age = 32
	location = "Dorne"

	fmt.Printf("%s (%d) of %s", name, age, location)
}
