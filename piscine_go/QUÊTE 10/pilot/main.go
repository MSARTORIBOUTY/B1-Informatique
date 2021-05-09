package main

import "fmt"

func main() {
	const AIRCRAFT1 = 1
	var donnie Pilot
	donnie.Name = "Donnie"
	donnie.Life = 100.0
	donnie.Age = 24
	donnie.Aircraft = AIRCRAFT1

	fmt.Println(donnie)
}

type Pilot struct {
	Name     string
	Life     int
	Age      int
	Aircraft int
}
