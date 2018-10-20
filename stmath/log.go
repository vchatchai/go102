package stmath

import (
	"fmt"
	"math"
)

func Log() {
	ln := math.Log(math.E)
	fmt.Printf("Ln(E) = %.4E\n", ln)

	log10 := math.Log10(-100)
	fmt.Printf("Log10(10) = %.4f\n", log10)

	log2 := math.Log2(2)
	fmt.Printf("Log2(2) = %.4f\n", log2)

	log_3_6 := Logarithm(3, 6)
	fmt.Printf("Log3(6) = %.4f\n", log_3_6)

}

func Logarithm(base, x float64) float64 {
	return math.Log(x) / math.Log(base)
}
