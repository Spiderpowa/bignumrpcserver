package bignumcalculator

import "math/big"

// BigNumberCalculator stores a mapping table for named number objects
type BigNumberCalculator struct {
	symbol map[string]big.Float
}

// New will create a BigNumber Calculator
func New() *BigNumberCalculator {
	calc := new(BigNumberCalculator)
	calc.symbol = make(map[string]big.Float)
	return calc
}

// Get the named object from the map. If name can be parsed as floating number, big.Float will be created.
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

// Set the named object in the map. Return false on error.
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

// Create a new named object and insert into the map.
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

// Delete the named object from the map
func (calc *BigNumberCalculator) Delete(name string) bool {
	_, exists := calc.symbol[name]
	if !exists {
		return false
	}
	delete(calc.symbol, name)
	return true
}
