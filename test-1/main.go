package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// To use this application you need to set the os argument in the terminal
	// For example, go run main.go 7
	// The output will be the solaneA000124 function results
	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("Solane's OEIS A000124 Application")
	fmt.Println(strings.Repeat("-", 35))

	inputScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input the number:")
inputInt:
	fmt.Printf("-> ")
	inputScanner.Scan()
	input, err := strconv.Atoi(strings.TrimSpace(inputScanner.Text()))
	if err != nil {
		fmt.Println("input must be an integer")
		goto inputInt
	}

	solaneResult := solaneA000124(input)

	output := make([]string, len(solaneResult))
	for k, result := range solaneResult {
		output[k] = fmt.Sprintf("%d", result)
	}

	fmt.Println(strings.Join(output, "-"))
}

func solaneA000124(n int) []int {
	// formula: n(n+1)/2 + 1
	formula := func(in int) int {
		return in*(in+1)/2 + 1
	}

	if n == 0 {
		return []int{formula(n)}
	}

	outputs := make([]int, n)
	for k := range outputs {
		outputs[k] = formula(k)
	}

	return outputs
}
