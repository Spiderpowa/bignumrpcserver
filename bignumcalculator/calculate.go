package bignumcalculator

import "math/big"

// Add x and y and return. second return value is success
func (calc *BigNumberCalculator) Add(x, y string) (big.Float, bool) {
	xval, yval := calc.Get(x), calc.Get(y)
	if xval == nil || yval == nil {
		return big.Float{}, false
	}

	return *new(big.Float).Add(xval, yval), true
}

// Sub x and y and return. second return value is success
func (calc *BigNumberCalculator) Sub(x, y string) (big.Float, bool) {
	xval, yval := calc.Get(x), calc.Get(y)
	if xval == nil || yval == nil {
		return big.Float{}, false
	}

	return *new(big.Float).Sub(xval, yval), true
}

// Mul x and y and return. second return value is success
func (calc *BigNumberCalculator) Mul(x, y string) (big.Float, bool) {
	xval, yval := calc.Get(x), calc.Get(y)
	if xval == nil || yval == nil {
		return big.Float{}, false
	}

	return *new(big.Float).Mul(xval, yval), true
}

// Div x and y and return. second return value is success
func (calc *BigNumberCalculator) Div(x, y string) (big.Float, bool) {
	xval, yval := calc.Get(x), calc.Get(y)
	if xval == nil || yval == nil {
		return big.Float{}, false
	}

	return *new(big.Float).Quo(xval, yval), true
}
