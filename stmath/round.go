package stmath

import (
	"fmt"
	"math"
)

var valA float64 = -3.55554444

func Round() {
	//Bad assumption on rounding
	//the number by casting it to
	//integer

	fmt.Printf("value %v\n", valA)
	fmt.Println(math.Copysign(1, valA))
	fmt.Println(math.Trunc(valA))
	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v \n", intVal)

	fRound := round(valA)
	fmt.Printf("Rounding by custom function: %v \n", fRound)

	fmt.Printf("Rounding by math.Round function: %v \n", math.Round(valA))
}

func round(x float64) float64 {
	t := math.Trunc(x)

	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}

	return t
}
