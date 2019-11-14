package main

import (
	"fmt"
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

	fmt.Printf("The expression of %s equals %f\n", first, calc(first))
	calcUI()
}

// / * - +

type parenthesis struct {
	left  string
	right string
}

type operator struct {
	symbol string
}

var bracket = parenthesis{left: "[", right: "]"}
var paren = parenthesis{left: "(", right: ")"}
var parens = []parenthesis{bracket, paren}

var division = operator{symbol: "/"}
var multiplication = operator{symbol: "*"}
var subtraction = operator{symbol: "-"}
var addition = operator{symbol: "+"}
var operators = []operator{division, multiplication, subtraction, addition}

func calc(exp string) float64 {
	f, err := strconv.ParseFloat(exp, 64)

	if err == nil {
		return f
	}

	cleanExp := strings.Split(strings.Replace(exp, " ", "", -1), "")

	for i := 0; i < len(parens); i++ {
		p := parens[i]
		if strings.Contains(exp, p.right) {

			var firstLeftSideIndex int
			var leftSideCount = 0
			var rightSideCount = 0
			for j := 0; j < len(cleanExp); j++ {
				// find the left side
				if cleanExp[j] == p.left {
					leftSideCount++

					if leftSideCount == 1 {
						firstLeftSideIndex = j
					}
				}

				// right the right side
				if cleanExp[j] == p.right {
					rightSideCount++

					if leftSideCount == rightSideCount {
						// start again
						return calc(strings.Join(cleanExp[firstLeftSideIndex+1:j], ""))
					}
				}
			}
		}
	}

	for i := 0; i < len(operators); i++ {
		o := operators[i]

		if strings.Contains(exp, o.symbol) {
			for j := 0; j < len(cleanExp); j++ {
				if cleanExp[j] == o.symbol {
					left := getLeftSideNum(cleanExp[0:j])
					right := getRightSideNum(cleanExp[j+1:])

					lStr := strings.Join(cleanExp[0:j-len(left)], "")
					r := mathOperation(o.symbol, getFloatValueFromArray(left), getFloatValueFromArray(right))
					rStr := strings.Join(cleanExp[j+len(right)+1:], "")

					// fmt.Printf("ls -%s\n", lStr)
					// fmt.Printf("r - %f\n", r)
					// fmt.Printf("rs - %s\n", rStr)

					return calc(fmt.Sprintf("%s%f%s", lStr, r, rStr))
				}
			}
		}
	}

	return -1
}

func getLeftSideNum(exp []string) []string {

	var startIndex int
	for i := len(exp) - 1; i >= 0; i-- {
		currChar := exp[i]

		if currChar == "." {
			continue
		}

		_, err := strconv.ParseFloat(currChar, 64)

		if err != nil {
			break
		}

		startIndex = i

	}

	return exp[startIndex:]
}

func getRightSideNum(exp []string) []string {

	// fmt.Printf("right - %s", strings.Join(exp, ""))
	var endIndex int
	var broke bool

	for i := 0; i < len(exp); i++ {
		endIndex = i

		currChar := exp[i]

		if currChar == "." || currChar == "0" {
			continue
		}

		_, err := strconv.ParseFloat(currChar, 64)

		if err != nil {
			broke = true
			break
		}

	}

	if !broke || endIndex == 0 {
		return exp[0:]
	}
	// fmt.Printf("right end %s - %d", strings.Join(exp[0:endIndex], ""), endIndex)
	return exp[0:endIndex]
}

func getFloatValueFromArray(numArray []string) float64 {
	numStr := strings.Join(numArray, "")
	// fmt.Printf("\n The String Num : %s \n", numStr)
	f, err := strconv.ParseFloat(numStr, 64)

	if err == nil {

		return f
	}

	return 0
}

func mathOperation(operator string, left float64, right float64) float64 {
	var result float64
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
