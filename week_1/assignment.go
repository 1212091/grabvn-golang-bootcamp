package main

import (
	"bufio"
	"fmt"
	"os"
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
