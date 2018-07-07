package main

import "github.com/Spiderpowa/bignumrpcserver/bignumcalculator"
import (
	"errors"
	"reflect"
)

// BigNumberHandler is a wrapper struct for BigNumberCalculator
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

// Create a new named object
func (t *BigNumberHandler) Create(args []interface{}, reply *string) error {
	params, err := handleArg(args)
	if err != nil {
		return err
	}
	if !t.calc.Create(params[0], params[1]) {
		return errors.New("Fail")
	}
	return nil
}

// Update a named object
func (t *BigNumberHandler) Update(args []interface{}, reply *string) error {
	params, err := handleArg(args)
	if err != nil {
		return err
	}
	if !t.calc.Set(params[0], params[1]) {
		return errors.New("Fail")
	}
	return nil
}

// Delete the named object
func (t *BigNumberHandler) Delete(args string, reply *string) error {
	if !t.calc.Delete(args) {
		return errors.New("Fail")
	}
	return nil
}

// Add x and y
func (t *BigNumberHandler) Add(args []interface{}, reply *string) error {
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

// Subtract x and y
func (t *BigNumberHandler) Subtract(args []interface{}, reply *string) error {
	params, err := handleArg(args)
	if err != nil {
		return err
	}
	ans, suc := t.calc.Sub(params[0], params[1])
	if !suc {
		return errors.New("Fail")
	}
	*reply = ans.String()
	return nil
}

// Multiply x and y
func (t *BigNumberHandler) Multiply(args []interface{}, reply *string) error {
	params, err := handleArg(args)
	if err != nil {
		return err
	}
	ans, suc := t.calc.Mul(params[0], params[1])
	if !suc {
		return errors.New("Fail")
	}
	*reply = ans.String()
	return nil
}

// Division x and y
func (t *BigNumberHandler) Division(args []interface{}, reply *string) error {
	params, err := handleArg(args)
	if err != nil {
		return err
	}
	ans, suc := t.calc.Div(params[0], params[1])
	if !suc {
		return errors.New("Fail")
	}
	*reply = ans.String()
	return nil
}

// CreateBigNumberHandler will return a new struct BigNumberHandler
func CreateBigNumberHandler() *BigNumberHandler {
	handler := new(BigNumberHandler)
	handler.calc = bignumcalculator.New()
	return handler
}
