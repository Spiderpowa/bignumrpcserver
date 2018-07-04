package bignumcalculator

import "testing"
import "math/big"

func makeTable(t *testing.T) *BigNumberCalculator {
    calc := New()
    numbers := []string{
        "0",
        "1.0",
        "-1.0",
        "2",
        "3.0",
        "-3.0",
        "4",
        "5.1",
        "3.14",
        "100000000000000000000000000000000",
        "999999999999999999999999999999999999",
        "-999999999999999999999999999999999999",
        "1000000000000000000000000000000000000",
        "-1000000000000000000000000000000000000",
    }
    for _, n := range numbers {
        calc.Create(n, n)
        if calc.Get(n) == nil {
            t.Errorf("Setting error %q", n)
        }
    }
    return calc
}

func TestAdd(t *testing.T) {
    calc := makeTable(t)
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

func TestSub(t *testing.T) {
    calc := makeTable(t)
    cases := []struct {
        x, y, ans string
    }{
        {"1.0", "1.0", "0"},
        {"1.0", "-1.0", "2.0"},
        {"0", "1.0", "-1.0"},
        {"1000000000000000000000000000000000000", "1.0", "999999999999999999999999999999999999"},
        {"1.0", "-999999999999999999999999999999999999", "1000000000000000000000000000000000000"},
        {"-1.0", "999999999999999999999999999999999999", "-1000000000000000000000000000000000000"},
    }
    for _, c := range cases {
        out, suc := calc.Sub(c.x, c.y)
        if !suc {
            t.Errorf("Sub Error %q %q", c.x, c.y)
            continue
        }
        ans, _ := new(big.Float).SetString(c.ans)
        if ans.Cmp(&out)!=0 {
            t.Errorf("Sub %q %q == %q, expect %q", c.x, c.y, out.String(), ans.String())
        }
    }
}

func TestMul(t *testing.T) {
    calc := makeTable(t)
    cases := []struct {
        x, y, ans string
    }{
        {"1.0", "1.0", "1"},
        {"1.0", "-1.0", "-1.0"},
        {"-3.0", "-1.0", "3.0"},
        {"0", "1.0", "0"},
        {"1000000000000000000000000000000000000", "1.0", "1000000000000000000000000000000000000"},
        {"1.0", "-999999999999999999999999999999999999", "-999999999999999999999999999999999999"},
        {"100000000000000000000000000000000", "100000000000000000000000000000000", "10000000000000000000000000000000000000000000000000000000000000000"},
    }
    for _, c := range cases {
        out, suc := calc.Mul(c.x, c.y)
        if !suc {
            t.Errorf("Mul Error %q %q", c.x, c.y)
            continue
        }
        ans, _ := new(big.Float).SetString(c.ans)
        if ans.Cmp(&out)!=0 {
            t.Errorf("Mul %q %q == %q, expect %q", c.x, c.y, out.String(), ans.String())
        }
    }
}

func TestDiv(t *testing.T) {
    calc := makeTable(t)
    cases := []struct {
        x, y, ans string
    }{
        {"4", "2", "2"},
        {"1.0", "-1.0", "-1.0"},
        {"-3.0", "-1.0", "3.0"},
        {"0", "1.0", "00"},
        {"1000000000000000000000000000000000000", "1.0", "1000000000000000000000000000000000000"},
        {"1000000000000000000000000000000000000", "1000000000000000000000000000000000000", "1"},
        {"1000000000000000000000000000000000000", "-1000000000000000000000000000000000000", "-1"},
    }
    for _, c := range cases {
        out, suc := calc.Div(c.x, c.y)
        if !suc {
            t.Errorf("Div Error %q %q", c.x, c.y)
            continue
        }
        ans, _ := new(big.Float).SetString(c.ans)
        if ans.Cmp(&out)!=0 {
            t.Errorf("Div %q %q == %q, expect %q", c.x, c.y, out.String(), ans.String())
        }
    }
}
