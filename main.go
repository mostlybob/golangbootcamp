package main

import "fmt"

func main() {

	//2.1 Variables & inferred typing
	var (
		name     string = "Prince Oberyn"
		age      int    = 32
		location string = "Dorne"
	)

	fmt.Printf("%s (%d) of %s", name, age, location)
}
