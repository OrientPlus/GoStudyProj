package main

import (
	"StudyProj/Cars"
	"fmt"
)

func main() {
	c := Cars.CreateCars(10, 4, 2, 4)
	for i, car := range c {
		fmt.Printf("===============\n[%d]\n", i)
		car.Print()
	}
}
