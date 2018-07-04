package bignumcalculator

import "math/big"

func (calc *BigNumberCalculator) Add(x, y string) (big.Float, bool) {
    xval, yval := calc.Get(x), calc.Get(y)
    if xval == nil || yval == nil {
        return big.Float{}, false
    }

    return *new(big.Float).Add(xval, yval), true
}
