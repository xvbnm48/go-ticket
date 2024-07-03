package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DivideOperation() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Divide")
	fmt.Print("Enter the first number: ")
	scanner.Scan()
	firstNumberStr := scanner.Text()
	firstNumber, _ := strconv.Atoi(firstNumberStr)
	fmt.Print("Enter the second number: ")
	scanner.Scan()
	secondNumberStr := scanner.Text()
	secondNumber, _ := strconv.Atoi(secondNumberStr)

	fmt.Println("Result:", divide(firstNumber, secondNumber))
}

func divide(firstNumber, secondNumber int) int {
	return firstNumber / secondNumber
}
