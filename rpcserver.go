package main

import "fmt"
import "github.com/Spiderpowa/bignumrpcserver/bignumcalculator"

func main() {
    calculator := bignumcalculator.New()
    calculator.Create("a", 5.5)
    fmt.Println(calculator)
}