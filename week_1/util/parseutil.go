package util

import (
	"errors"
	"strconv"
	"strings"
)

func ParseInput(input string) (float64, float64, string, error) {
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

	a, err := parseNumber(expr, 0)

	if err != nil {
		return 0, 0, "", err
	}

	b, err := parseNumber(expr, 1)

	return a, b, input[operatorIndex : operatorIndex+1], err
}

func parseNumber(expr []string, exprIndex int) (float64, error) {
	first := strings.Replace(expr[exprIndex], " ", "", -1)

	a, err := strconv.ParseFloat(first, 64)

	if err != nil {
		err := errors.New("Invalid number")
		return 0, err
	}
	return a, err
}
