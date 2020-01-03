package main

import (
	"fmt"
	"strconv"
)

func main() {
	stringIn := "0"
	res := diffWaysToCompute(stringIn)

	fmt.Println("ANSWER: ", res)
}

func Calculate(na int, x string, nb int) int {
	switch x {
	case "+":
		return na + nb
	case "-":
		return na - nb
	case "*":
		return na * nb
	default:
		panic("Input error")
	}
}

func RecurCount(numbers []string, operator []string, iStart int, iEnd int) []int {
	result := []int{}

	if iStart == iEnd {
		intVal, _ := strconv.Atoi(numbers[iStart])
		result = append(result, intVal)
	} else {
		for iNow := iStart; iNow < iEnd; iNow++ {

			leftSide := RecurCount(numbers, operator, iStart, iNow)
			rightSide := RecurCount(numbers, operator, iNow+1, iEnd)

			for _, left := range leftSide {
				for _, right := range rightSide {
					result = append(result, Calculate(left, operator[iNow], right))
				}
			}
		}
	}
	return result
}

func splitInput(input string)(output []string) {
	tmp := ""

	for i := range input {
		if input[i] != '+' && input[i] != '-' && input[i] != '*' {
			tmp += string(input[i])
		} else {
			output = append(output, tmp)
			tmp = ""
			output = append(output, string(input[i]))
		}
		if i+1 == len(input) {
			output = append(output, tmp)
		}
	}
	return output
}

func diffWaysToCompute(input string) []int {
	slice := splitInput(input)
	numbers := []string{}
	operators := []string{}

	for _, i := range slice {
		if i == "+" || i == "-" || i == "*" {
			operators = append(operators, i)
		} else {
			numbers = append(numbers, i)
		}
	}
	if len(operators) == 0 && len(numbers) == 1 {
		num, _ := strconv.Atoi(numbers[0])
		answer := []int{num}
		return answer
	}
	return RecurCount(numbers, operators, 0, len(numbers)-1)
}
