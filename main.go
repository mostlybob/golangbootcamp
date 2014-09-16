package main

import "fmt"

func main() {

	//2.1 Variables & inferred typing
	action := func() {
		name, location := "Prince Oberyn", "Dorne"
		age := 32

		fmt.Printf("%s (%d) of %s", name, age, location)
	}

	action()
}
