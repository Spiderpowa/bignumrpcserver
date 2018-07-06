package main

import "github.com/Spiderpowa/bignumrpcserver/bignumcalculator"
import (
        "reflect"
        "errors"
)

type BigNumberHandler struct {
    calc *bignumcalculator.BigNumberCalculator
}

func handleArg(args []interface{}) ([]string, error) {
    if len(args) != 2 {
        return nil, errors.New("Expected param of length 2")
    }
    params := make([]string, 2)
    for i, value := range args {
        if reflect.TypeOf(value).Kind() != reflect.String {
            return nil, errors.New("Unknown type in param")
        }
        params[i] = value.(string)
    }
    return params, nil
}

func (t* BigNumberHandler) Create(args []interface{}, reply *string) error {
    params, err := handleArg(args)
    if err != nil {
        return err
    }
    if !t.calc.Create(params[0], params[1]) {
        return errors.New("Fail")
    }
    return nil
}

func (t* BigNumberHandler) Add(args []interface{}, reply *string) error {
    params, err := handleArg(args)
    if err != nil {
        return err
    }
    ans, suc := t.calc.Add(params[0], params[1])
    if !suc {
        return errors.New("Fail")
    }
    *reply = ans.String()
    return nil
}


func CreateBigNumberHandler() *BigNumberHandler {
    handler := new(BigNumberHandler)
    handler.calc = bignumcalculator.New()
    return handler
}
