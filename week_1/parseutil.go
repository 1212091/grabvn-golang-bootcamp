package main

import (
	"errors"
	"strconv"
	"strings"
)

func parseInput(input string) (float64, float64, string, error) {
	operatorIndex := strings.IndexAny(input, "+-*/")

	if operatorIndex == -1 {
		err := errors.New("Missing Operator")
		return 0, 0, "", err
	}

	expr := strings.Split(input, input[operatorIndex:operatorIndex+1])
	if len(expr) != 2 {
		err := errors.New("The number argument of input is invalid")
		return 0, 0, "", err
	}

	first := strings.Replace(expr[0], " ", "", -1)

	a, err := strconv.ParseFloat(first, 64)

	if err != nil {
		err := errors.New("Invalid first number")
		return 0, 0, "", err
	}

	second := strings.Replace(expr[1], " ", "", -1)

	b, err := strconv.ParseFloat(second, 64)
	if err != nil {
		err := errors.New("Invalid second number")
		return 0, 0, "", err
	}
	return a, b, input[operatorIndex : operatorIndex+1], nil
}
