package ovo

import (
	"errors"
	"fmt"
)

var (
	// ErrInternal is general internal error.
	ErrInternal = errors.New("internal error")
)

func errRequiredField(str string) error {
	return fmt.Errorf("required field %s", str)
}

func errInvalidValueField(str string) error {
	return fmt.Errorf("invalid %s value", str)
}

func errGTField(str, value string) error {
	return fmt.Errorf("field %s must be greater than %s", str, value)
}

func errLTField(str, value string) error {
	return fmt.Errorf("field %s must be lower than %s", str, value)
}

var ovoErr = map[string]error{
	"13": errors.New("invalid amount"),
	"14": errors.New("invalid mobile number/OVO ID"),
	"17": errors.New("transaction declined/cancelled"),
	"25": errors.New("transaction not found"),
	"26": errors.New("transaction failed"),
	"40": errors.New("transaction failed"),
	"54": errors.New("transaction expired"),
	"58": errors.New("transaction not allowed"),
	"63": errors.New("authentication failed"),
	"68": errors.New("transaction pending/timeout"),
	"73": errors.New("transaction has been reversed"),
	"94": errors.New("duplicate merchant invoice or reference number"),
	"96": errors.New("invalid processing code"),
	"ER": errors.New("internal OVO failure"),
	"EB": errors.New("tid/mid not registered"),
	"BR": errors.New("invalid request format"),
}
