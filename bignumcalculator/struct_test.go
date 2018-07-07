package bignumcalculator

import "testing"
import "math/big"

func TestCreate(t *testing.T) {
    cases := []struct {
        in []string
        out []bool
    }{
        {[]string{"Name1"}, []bool{true}},
        {[]string{"Name1", "Name2"}, []bool{true, true}},
        {[]string{"Name1", "Name1"}, []bool{true, false}},
    }
    for _, c := range cases {
        calc := New()
        for idx := range c.in {
            out := calc.Create(c.in[idx], "1.0")
            if out != c.out[idx] {
                t.Errorf("Create[%d](%q) == %t, expect %t", idx, c.in[idx], out, c.out[idx])
            }
        }
    }
}

func TestCreateVal(t *testing.T) {
    cases := []struct {
        in string
        out bool
    }{
        {"1.0", true},
        {"abcd", false},
        {"0x100", true},
        {".3", true},
        {"1.5a", false},
    }
    for _, c := range cases {
        calc := New()
        out := calc.Create("N", c.in)
        if out != c.out {
            t.Errorf("Create (%q) == %t, expect %t", c.in, out, c.out)
        }
        
    }
}

func TestGet(t *testing.T) {
    calc := New()
    calc.Create("N", "1.0")
    cases := []struct {
        in string
        out bool
    }{
        {"N", true},
        {"A", false},
        {"33.10", true},
        {"42", true},
        {"1e5", true},
    }
    for _, c := range cases {
        out := calc.Get(c.in) != nil
        if out != c.out {
            t.Errorf("Create (%q) == %t, expect %t", c.in, out, c.out)
        }
        
    }
}

func TestSet(t *testing.T) {
    calc := New()
    calc.Create("N", "1.0")
    cases := []struct {
        in, val string
        out bool
    }{
        {"N", "3.0", true},
        {"A", "5", false},
        {"N", "4.0", true},
        {"N", "WRONG", false},
    }
    for _, c := range cases {
        out := calc.Set(c.in, c.val)
        if out != c.out {
            t.Errorf("Set (%q, %q) == %t, expect %t", c.in, c.val, out, c.out)
        }
        if out {
            val := calc.Get(c.in)
            exp, _ := new(big.Float).SetString(c.val)
            if val.Cmp(exp) != 0 {
                t.Errorf("Set Fail (%q, %q) == %q", c.in, c.val, val.String())
            }
        }
    }
}

func TestDelete(t *testing.T) {
    calc := New()
    calc.Create("A", "1.0")
    calc.Create("B", "1.0")
    cases := []struct {
        in string
        del bool
        out bool
    }{
        {"A", true, true},
        {"B", false, true},
        {"C", true, false},
    }
    for _, c := range cases {
        if !c.del {
            continue
        }
        out := calc.Delete(c.in)
        if out != c.out {
            t.Errorf("Delete (%q) == %t, expect %t", c.in, out, c.out)
        }
        if out {
            if calc.Get(c.in) != nil {
                t.Errorf("Delete (%q) but still can be found", c.in)
            }
        }
    }
    for _, c := range cases {
        if c.del {
            continue
        }
        if calc.Get(c.in) == nil {
            t.Errorf("(%q) is not deleted but cannot be found", c.in)
        }
    }
}
