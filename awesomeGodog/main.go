package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return factorial(n-1) * n
}

func formatNiceFactorial(n int, expanded bool) string {
	var expansion string

	if expanded {
		expansion += " = 1"

		for i := 2; i <= n; i++ {
			expansion += fmt.Sprintf(" * %d", i)
		}
	}

	return fmt.Sprintf("%v!%v = %v", n, expansion, factorial(n))
}

// Example test cases:
// factorial(5) => 120
// factorial(0) => 1
// factorial(1) => 1
// formatNiceFactorial(5, false) => "5! = 120"
// formatNiceFactorial(5, true) => "5! = 1 * 2 * 3 * 4 * 5 = 120"
// formatNiceFactorial(1, true) => "1! = 1 = 1"
// formatNiceFactorial(0, true) => "0! = 1 = 1"
// formatNiceFactorial(3, false) => "3! = 6"

func main() {
	fmt.Println(formatNiceFactorial(5, true))
}
