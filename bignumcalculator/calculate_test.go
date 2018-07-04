package bignumcalculator

import "testing"
import "math/big"

func TestAdd(t *testing.T) {
    calc := New()
    numbers := []string{"1.0", "-1.0", "3.0", "5.1", "3.14", "999999999999999999999999999999999999"}
    for _, n := range numbers {
        calc.Create(n, n)
        if calc.Get(n) == nil {
            t.Errorf("Setting error %q", n)
        }
    }
    cases := []struct {
        x, y, sum string
    }{
        {"1.0", "1.0", "2.0"},
        {"1.0", "-1.0", "0"},
        {"1.0", "999999999999999999999999999999999999", "1000000000000000000000000000000000000"},
    }
    for _, c := range cases {
        out, suc := calc.Add(c.x, c.y)
        if !suc {
            t.Errorf("Add Error %q %q", c.x, c.y)
            continue
        }
        sum, _ := new(big.Float).SetString(c.sum)
        if sum.Cmp(&out)!=0 {
            t.Errorf("Add %q %q == %q, expect %q", c.x, c.y, out.String(), sum.String())
        }
    }
}
