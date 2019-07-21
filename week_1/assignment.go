package main

import (
	"bufio"
	"fmt"
	"os"

	"./util"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	for scanner.Scan() {
		input := scanner.Text()
		a, b, op, err := util.ParseInput(input)

		if err != nil {
			fmt.Println(err)
			printMark()
			continue
		}

		calculate(a, b, op)
		printMark()
	}
}

func calculate(a float64, b float64, op string) {
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
}

func printMark() {
	fmt.Print("> ")
}
