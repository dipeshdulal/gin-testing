package main

import "errors"

var (
	ErrNumberPositive = errors.New("number should be greater than zero")
)
