package bignumcalculator

import "math/big"

type BigNumberCalculator struct {
	symbol map[string]big.Float
}

func New() *BigNumberCalculator {
	calc := new(BigNumberCalculator)
	calc.symbol = make(map[string]big.Float)
	return calc
}

func (calc *BigNumberCalculator) Get(name string) *big.Float {
	val, exists := calc.symbol[name]
	if !exists {
		constVal, suc := new(big.Float).SetString(name)
		if suc {
			val = *constVal
			exists = true
		}
	}
	if !exists {
		return nil
	}
	return &val
}

func (calc *BigNumberCalculator) Set(name, val string) bool {
	_, exists := calc.symbol[name]
	if !exists {
		return false
	}
	value, suc := new(big.Float).SetString(val)
	if suc {
		calc.symbol[name] = *value
	}
	return suc
}

func (calc *BigNumberCalculator) Create(name, val string) bool {
	_, exists := calc.symbol[name]
	if exists {
		return false
	}
	value, suc := new(big.Float).SetString(val)
	if suc {
		calc.symbol[name] = *value
	}
	return suc
}

func (calc *BigNumberCalculator) Delete(name string) bool {
	_, exists := calc.symbol[name]
	if !exists {
		return false
	}
	delete(calc.symbol, name)
	return true
}
