package stmath

import (
	"fmt"
	"math/cmplx"
)

func Complex() {
	a := complex(2, 3)
	b := complex(6, 4)

	fmt.Printf("Real part: %f\n", real(a))
	fmt.Printf("Complex part: %f\n", imag(a))

	fmt.Printf(" a: %v\n", a)
	fmt.Printf(" b: %v\n", b)
	c := a - b
	fmt.Printf("Difference : %v \n", c)
	c = a + b
	fmt.Printf("Sum %v\n", c)
	c = a * b
	fmt.Printf("Product : %v\n", c)
	c = a / b
	fmt.Printf("Devide : %v\n", c)

	conjugate := cmplx.Conj(a)
	fmt.Printf("Complex number a's conjugate : %v\n ", conjugate)

	cos := cmplx.Cos(b)
	fmt.Printf("Cosine of b: %v\n", cos)
}
