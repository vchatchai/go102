package stmath

import (
	"fmt"
	"math"
)

type Radian float64
type Degree float64

func (rad Radian) toDegrees() Degree {
	return Degree(float64(rad) * (180.0 / math.Pi))
}

func (rad Radian) toFloat64() float64 {
	return float64(rad)
}

func (deg Degree) toRadian() Radian {
	return Radian(float64(deg) * math.Pi / 180.0)
}

func (deg Degree) toFloat64() float64 {
	return float64(deg)
}

func Radians() {
	val := radiansToDegrees(1)
	fmt.Printf("One radian is :%.4f degreee\n", val)

	val2 := degreesToRadian(val)
	fmt.Printf("%.4f degrees is %.4f degrees\n", val, val2)

	val = Radian(1).toDegrees().toFloat64()
	fmt.Printf("Degree: %.4f degree\n", val)

	val = Degree(val).toRadian().toFloat64()
	fmt.Printf("Rad: %.4f radians\n", val)
}

func radiansToDegrees(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}

func degreesToRadian(rad float64) float64 {
	return rad * (180.0 / math.Pi)
}
