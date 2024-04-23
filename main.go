package main

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
)

func main() {
	//Logo
	fmt.Println(`
 __  __          ____   _                           _          
|  \/  |        |  _ \ | |                         | |          TM
| \  / | _ __   | |_) || |__    __ _  _ __   _   _ | | __  __ _ 
| |\/| ||  __|  |  _ < |  _ \  / _  ||  _ \ | | | || |/ / / _  |
| |  | || |  _  | |_) || | | || (_| || | | || |_| ||   < | (_| |
|_|  |_||_| (_) |____/ |_| |_| \__ _||_| |_| \__ _||_|\_\ \____|

Limits Simulator | Version 1.0.0
`)

	// Notes
	fmt.Println(`
Notes:
	- The function variable should be denoted by 'x'.
	- Trigonometric functions (e.g., sin, cos, tan) are not supported.
	- The value of X cannot be set to infinity (e.g., x->infinity is not supported).
	- Only these arithmetic operators are supported: -, *, /, ** or ^, %
	  (** and ^ refers to "take to the power of", for example 2**2 or 2^2 is equal to 4.).
	  `)

	// Variables
	var function string
	var x, point, evaluate float64

	// User Inputs
	fmt.Println("Enter the function:")
	fmt.Scanln(&function)
	fmt.Println("Value of X (x->?):")
	fmt.Scan(&x)
	fmt.Println("Range around X (e.g., if you enter 2, it will use 2 -> -2):")
	fmt.Scan(&point)
	fmt.Println("Step size for evaluation:")
	fmt.Scan(&evaluate)

	// Replace "^" with "**" for exponentiation
	function = strings.ReplaceAll(function, "^", "**")

	expression, err := govaluate.NewEvaluableExpression(function)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	// Calculations
	calculate(expression, x, point, evaluate)
}

func calculate(expression *govaluate.EvaluableExpression, x, point, evaluate float64) {
	parameters := make(map[string]interface{})

	for i := -point; i <= 0; i += evaluate {
		xValue := x + i
		parameters["x"] = xValue
		// Evaluate the expression with the new x value
		result, err := expression.Evaluate(parameters)
		if err != nil {
			fmt.Println("Error evaluating expression:", err)
			return
		}
		fmt.Printf("⬇️ %.2f -> %.2f\n", xValue, result)
	}

	fmt.Println("🔽")
	xValue := x
	parameters["x"] = xValue
	// Evaluate the expression with the new x value
	result, err := expression.Evaluate(parameters)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}
	fmt.Printf("\n%.2f -> %.2f\n\n", xValue, result)
	defer func() {
		fmt.Printf("\n\n\n Limit of %s as (x -> %.2f) approaches %.2f\n\n", expression, xValue, result)
	}()
	fmt.Println("🔼")
	for i := evaluate; i <= point; i += evaluate {
		xValue = x + i
		parameters["x"] = xValue
		// Evaluate the expression with the new x value
		result, err := expression.Evaluate(parameters)
		if err != nil {
			fmt.Println("Error evaluating expression:", err)
			return
		}
		fmt.Printf("⬆️ %.2f -> %.2f\n", xValue, result)
	}
}
