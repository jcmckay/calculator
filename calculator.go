package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	calcUI()
}

func calcUI() {
	fmt.Print("Enter string to calculate: ")
	var first string
	fmt.Scanln(&first)

	if first == "exit" {
		return
	}

	fmt.Printf("The expression of %s equals %d\n", first, calculator(first))
	calcUI()
}

func calculator(expression string) int64 {
	formatted := strings.Split(strings.Replace(expression, " ", "", -1), "")
	return leftRightCalc(formatted, 0, "")
}

func leftRightCalc(exp []string, result int64, operator string) int64 {

	if len(exp) == 0 {
		return result
	}

	if isInt(exp[0]) == -1 {
		return leftRightCalc(exp[1:], result, exp[0])
	} else {
		// look forward to see the full number
		var currNumLength = 1
		if len(exp) > 1 {
			for j := 1; j <= len(exp); j++ {
				currNumLength = j

				if j == len(exp) || isInt(exp[j]) == -1 {
					break
				}

			}
		}

		num, e := strconv.ParseInt(strings.Join(exp[0:currNumLength], ""), 0, 64)

		if e == nil {
			result = mathOperation(operator, result, num)
		} else {
			log.Fatal("This should always pass")
		}

		return leftRightCalc(exp[currNumLength:], result, "")
	}

}

func mathOperation(operator string, left int64, right int64) int64 {
	var result int64
	switch operator {
	case "+":
		result = left + right
	case "-":
		result = left - right
	case "*":
		result = left * right
	case "/":
		result = left / right
	default:
		result = left + right
	}

	return result
}

func isInt(s string) int64 {
	i, err := strconv.ParseInt(s, 0, 64)

	if err != nil {
		return -1
	}

	return i
}
