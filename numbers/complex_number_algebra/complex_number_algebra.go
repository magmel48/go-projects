/**
Show addition, multiplication, negation, and inversion of complex numbers in separate functions.
(Subtraction and division operations can be made with pairs of these operations.)
Print the results for each operation tested.

NOTE definitely everything is already written in Go (complex64)
*/
package complex_number_algebra

import "fmt"

type complexNumber struct {
	realPart      float64
	imaginaryPart float64
}

func cmplx(r float64, i float64) complexNumber {
	return complexNumber{r, i}
}

func plus(a complexNumber, b complexNumber) (out complexNumber) {
	return cmplx(a.realPart+b.realPart, a.imaginaryPart+b.imaginaryPart)
}

func minus(a complexNumber, b complexNumber) (out complexNumber) {
	return cmplx(a.realPart-b.realPart, a.imaginaryPart-b.imaginaryPart)
}

func multiplication(a complexNumber, b complexNumber) (out complexNumber) {
	return cmplx(
		a.realPart*b.realPart-a.imaginaryPart*b.imaginaryPart,
		a.imaginaryPart*b.realPart+a.realPart*b.imaginaryPart)
}

func division(a complexNumber, b complexNumber) (out complexNumber) {
	multiplier := 1 / (b.realPart*b.realPart + b.imaginaryPart*b.imaginaryPart)

	return cmplx(
		multiplier*(a.realPart*b.realPart+a.imaginaryPart*b.imaginaryPart),
		multiplier*(a.imaginaryPart*b.realPart-a.realPart*b.imaginaryPart))
}

func ComplexNumberAlgebra() {
	a := cmplx(1, 2)
	b := cmplx(3, 4)

	fmt.Println("plus:", plus(a, b))
	fmt.Println("minus:", minus(a, b))
	fmt.Println("multiplication:", multiplication(a, b))
	fmt.Println("division:", division(a, b))
}
