package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		input := scanner.Text()
		a, b, op, err := parseInput(input)

		if err != nil {
			fmt.Println(err)
			fmt.Print("> ")
			continue
		}

		switch op {
		case "+":
			fmt.Println(a, "+", b, "=", a+b)
		case "-":
			fmt.Println(a, "-", b, "=", a-b)
		case "*":
			fmt.Println(a, "*", b, "=", a*b)
		case "/":
			if b == 0 {
				fmt.Println("Cannot divide to 0")
			} else {
				fmt.Println(a, "/", b, "=", a/b)
			}
		default:
			fmt.Println("Invalid Operator")
		}
		fmt.Print("> ")
	}
}

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
