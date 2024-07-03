package main

import (
	"bufio"
	"fmt"
	"github.com/xvbnm48/golang-terminal-learn/internal/calculator"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("=============")
		fmt.Println("CALCULATOR")
		fmt.Println("=============")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Print("Choose the operation: ")
		scanner.Scan()
		option := scanner.Text()
		switch option {
		case "1":
			calculator.AddOperation()
			break
		case "2":
			calculator.SubstractOperation()
			break
		case "3":
			calculator.MultiplyOperation()
			break
		case "4":
			calculator.DivideOperation()
			break
		case "5":
			fmt.Println("Exit")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
		}
	}
}
