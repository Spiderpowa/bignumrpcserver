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
            out := calc.Create(c.in[idx], 1.0)
            if out != c.out[idx] {
                t.Errorf("Create[%d](%q) == %t, expect %t", idx, c.in[idx], out, c.out[idx])
            }
        }
    }
}