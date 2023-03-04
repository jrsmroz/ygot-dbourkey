package util

import (
	"errors"
	"fmt"
	"strconv"
)

type ParseYangIntegerError struct {
	Func string // the failing function (ParseYangUint)
	Num  string // the input
	Err  error  // the reason the conversion failed
}

func (e *ParseYangIntegerError) Error() string {
	return fmt.Sprintf("util.%s: parsing %v: %s", e.Func, e.Num, e.Err.Error())
}

func (e *ParseYangIntegerError) Unwrap() error { return e.Err }

// ErrYangIntegerSyntax indicates that a value does not have the right syntax for the supported yang integer types.
var ErrYangIntegerSyntax = errors.New("invalid syntax")

func syntaxError(fn string, num string) error {
	return &ParseYangIntegerError{
		Func: fn,
		Num:  num,
		Err:  ErrYangIntegerSyntax,
	}
}

func strconvError(fn string, num string, err error) error {
	return &ParseYangIntegerError{
		Func: fn,
		Num:  num,
		Err:  err,
	}
}

func ParseYangUint(s string, bitSize int) (uint64, error) {
	const fnParseYangUint = "ParseYangUint"
	if s == "" {
		return 0, syntaxError(fnParseYangUint, s)
	}

	if len(s) > 1 && s[0] == '+' {
		s = s[1:]
	}

	var base int
	switch {
	case len(s) > 2 && s[:2] == "0x":
		s = s[2:]
		base = 16
	default:
		base = 10
	}

	val, err := strconv.ParseUint(s, base, bitSize)
	if err != nil {
		return 0, strconvError(fnParseYangUint, s, err)
	}

	return val, nil
}

func ParseYangInt(s string, bitSize int) (int64, error) {
	const fnParseYangUint = "ParseYangInt"
	if s == "" {
		return 0, syntaxError(fnParseYangUint, s)
	}

	sign := "+"
	if len(s) > 1 && s[0] == '+' {
		s = s[1:]
	}
	if len(s) > 1 && s[0] == '-' {
		s = s[1:]
		sign = "-"
	}

	var base int
	switch {
	case len(s) > 2 && s[:2] == "0x":
		s = s[2:]
		base = 16
	default:
		base = 10
	}

	val, err := strconv.ParseInt(sign+s, base, bitSize)
	if err != nil {
		return 0, strconvError(fnParseYangUint, s, err)
	}

	return val, nil
}
