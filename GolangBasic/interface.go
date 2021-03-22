package main

import "fmt"

type Vehicle interface {
	Introduce() string
	Accelerate(factor int) int
}

type Plane struct {
	name     string
	speed    int
	wingspan int
}

type Car struct {
	name       string
	speed      int
	numOfTires int
}

func (car *Car) Introduce() string {
	return fmt.Sprintf("%v, %v, %v", car.name, car.speed, car.numOfTires)
}

func (car *Car) Accelerate(factor int) int {
	return factor * car.speed
}

func (plane *Plane) Introduce() string {
	return fmt.Sprintf("%v - %v  - %v", plane.name, plane.speed, plane.wingspan)
}

func (plane *Plane) Accelerate(factor int) int {
	return factor * plane.speed
}

func main() {
	var vechicles []Vehicle
	car := Car{
		name:       "BMW",
		speed:      100,
		numOfTires: 4,
	}
	plane := Plane{
		name:     "Boeing",
		speed:    500,
		wingspan: 100,
	}

	vechicles = append(vechicles, &car)
	vechicles = append(vechicles, &plane)
	fmt.Println(vechicles)
}
