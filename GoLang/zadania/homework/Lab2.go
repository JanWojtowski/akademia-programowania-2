package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	var number1 int
	var number2 int

	fmt.Print("Type first number: ")
	fmt.Scan(&number1)
	fmt.Println("Your firstnumber is:", number1)

	fmt.Print("Type second number: ")
	fmt.Scan(&number2)
	fmt.Println("Your second number is:", number2)

	var operation int

	fmt.Println("Operation: \n 1. Add \n 2. Subtract \n 3. Multiply \n 4. Divide")
	fmt.Print("Choose operation: ")
	fmt.Scan(&operation)

	operations(operation, number1, number2)
}

func operations(operation int, number1 int, number2 int) int {
	output, err := os.OpenFile("Results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := output.Close(); err != nil {
			log.Panicln(err)
		}
	}()

	a := big.NewInt(int64(number1))
	b := big.NewInt(int64(number2))
	sum := big.NewInt(0)

	switch operation {
	case 1:
		sum := sum.Add(a, b)
		_, err := output.WriteString(fmt.Sprintf("%d + %d = %d \n", a, b, sum))
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(fmt.Sprintf("Operation result: %d + %d = %d", number1, number2, sum))
		return int(sum.Int64())
	case 2:
		sum.Sub(a, b)
		writer := bufio.NewWriter(output)
		if _, err := writer.WriteString(fmt.Sprintf("%d - %d = %d \n", a, b, sum)); err != nil {
			log.Panicln(err)
		}
		err := writer.Flush()
		if err != nil {
			return 0
		}
		fmt.Printf(fmt.Sprintf("Operation result: %d - %d = %d", number1, number2, sum))
		return int(sum.Int64())
	case 3:
		sum := sum.Mul(a, b)
		_, err := output.WriteString(fmt.Sprintf("%d * %d = %d \n", a, b, sum))
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(fmt.Sprintf("Operation result: %d * %d = %d", number1, number2, sum))
		return int(sum.Int64())
	case 4:
		if number2 == 0 {
			fmt.Println("You can divide by 0")
			fmt.Println("Choose new second number:")
			fmt.Scan(&number2)
			operations(operation, number1, number2)
		} else {
			sum.Div(a, b)
			_, err := output.WriteString(fmt.Sprintf("%d / %d = %d \n", a, b, sum))
			if err != nil {
				log.Panicln(err)
			}
			fmt.Println(fmt.Sprintf("Operation result: %d / %d =  %d", number1, number2, sum))
			return int(sum.Int64())
		}
	default:
		return 0
	}
	return 0
}
