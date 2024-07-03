package calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func AddOperation() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("\n==============================")
	fmt.Println("       Addition Operation")
	fmt.Println("==============================")
	fmt.Print("Enter the first number: ")
	scanner.Scan()
	firstNumberStr := scanner.Text()
	firstNumber, _ := strconv.Atoi(firstNumberStr)
	fmt.Print("Enter the second number: ")
	scanner.Scan()
	secondNumberStr := scanner.Text()
	secondNumber, _ := strconv.Atoi(secondNumberStr)

	fmt.Println("Result:", add(firstNumber, secondNumber))

}

func add(firstNumber, secondNumber int) int {
	return firstNumber + secondNumber
}
