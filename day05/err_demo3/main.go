package main

import (
	"fmt"
	"math"
)

func main() {
	radius := -0.12
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("radius  %.2f is less than zero", err.radius)
			return
		}
		fmt.Println(err)
	}
	fmt.Printf("area is : %.2f", area)

	for i := 0; i < 9; i++ {
		fmt.Println(i)
	}

}

type areaError struct {
	radius float64
	err    string
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %.2f, %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{radius, "radius is negative"}
	}
	return math.Pi * math.Pow(radius, 2), nil
}
