package bignumcalculator

type BigNumberCalculator struct{
    symbol map[string]float64
}

func New() *BigNumberCalculator {
    calc := new(BigNumberCalculator)
    calc.symbol = make(map[string]float64)
    return calc
}

func (calc *BigNumberCalculator) Create(name string, val float64) bool {
    _, exists := calc.symbol[name]
    if exists {
        return false
    }
    calc.symbol[name] = val
    return true
}