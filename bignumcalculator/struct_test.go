package bignumcalculator

import "testing"

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
    }
    for _, c := range cases {
        out := calc.Get(c.in) != nil
        if out != c.out {
            t.Errorf("Create (%q) == %t, expect %t", c.in, out, c.out)
        }
        
    }
}

