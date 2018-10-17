package stfmt

import "fmt"

var integer int64 = 32500
var floatNum float64 = 22000.456

func Format() {
	//Common way how to print the decimal
	//num
	fmt.Printf("decimal: %d \n", integer)
	//Always show the sign
	fmt.Printf("decimal with sign:%+d \n", integer)
	//print in other base X -16, 0-8, b-2, d -10
	fmt.Printf("X:%X \n", integer)
	fmt.Printf("#X:%#X \n", integer)

	//Padding with  leading zeros
	fmt.Printf("padding:%010d \n", integer)

	//Left padding with spaces
	fmt.Printf("left padding:% 10d \n", integer)

	//Right padding
	fmt.Printf("right padding:% -10d \n", integer)

	//Print floatingpoint number
	fmt.Printf("floating-point:%f \n", floatNum)

	// Floating-point number  with limited precision = 5
	fmt.Printf("precision 5:%.5f \n", floatNum)

	//Floating-point number in scientitic notation
	fmt.Printf("scientitic notation e:%e  \n", floatNum)

	//Floating-point number %e for large exponents or %f  otherwise
	fmt.Printf("g:%g \n", floatNum)
}
