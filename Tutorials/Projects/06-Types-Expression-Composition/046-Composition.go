//046-Composition.go
package main

import "fmt"

func main() {

	pickup := Vehicle {
		NumberOfWheels: 4,
		NumberOfPassengers: 6,
	}

	unimog := Car {
		Make: "Mercedes Benz",
		Model: "U 5000",
		Year: 2020,
		IsElectric: true,
		IsHybrid: false,
		Vehicle: pickup,
	}

	unimog.show()
}

type Vehicle struct {
	NumberOfWheels int
	NumberOfPassengers int
} 

type Car struct {
	Make string
	Model string
	Year int
	IsElectric bool
	IsHybrid bool
	Vehicle Vehicle
}

func (v Vehicle) showDetails() {
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
	fmt.Println("Number of passengers:", v.NumberOfWheels)
}

func (c Car) show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is elecric:", c.IsElectric)
	fmt.Println("Is hybrid:", c.IsHybrid)
	fmt.Println("Number of wheels:", c.Vehicle.NumberOfWheels)
	fmt.Println("Number of passengers:", c.Vehicle.NumberOfPassengers)
}
