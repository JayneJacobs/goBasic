package interfaces

import "fmt"

// CarI allows access to Car methods
type CarI interface {
	DescribeCar(c *Car) *Car
}

// Car has attributes to the car
type Car struct {
	wh  string
	sts int
	drs int
	t   bool
	bs  string
}

// DescribeCar defines the car
func (c *Car) DescribeCar(car *Car) *Car {
	if c.t {
		fmt.Printf("This is a car: %v\n", &c)
		return c
	}

	return c
}

// ImplementCar assigns values to a sedan
func ImplementCar() {
	sedan := &Car{"Leather", 7, 4, true, "large"}
	var thisCar Car
	myCar := sedan.DescribeCar(&thisCar)
	fmt.Printf("This is my car description: %v\n", *myCar)
}

//And interface declares what methods an object has as a set of behaviors.
//This is an example A car drives and parks. A type that has these methods implements the car interface
//A car has doors that open. A car has a back seat, a car has a trunk. The trunk lid opens....
//So a type that implements these behaviours is said to implement the car interface for example a sedan would implement the car interface.
