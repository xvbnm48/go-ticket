package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SubstractOperation() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Subtract")
	fmt.Print("Enter the first number: ")
	scanner.Scan()
	firstNumberStr := scanner.Text()
	firstNumber, _ := strconv.Atoi(firstNumberStr)
	fmt.Print("Enter the second number: ")
	scanner.Scan()
	secondNumberStr := scanner.Text()
	secondNumber, _ := strconv.Atoi(secondNumberStr)

	fmt.Println("Result:", subtract(firstNumber, secondNumber))
}

func subtract(firstNumber, secondNumber int) int {
	return firstNumber - secondNumber
}
