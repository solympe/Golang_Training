package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stringIn := "11+1"
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

func diffWaysToCompute(input string) []int {
	slice := strings.Split(input, "")
	alert := false




	numbers := []string{}
	operators := []string{}

	for i := range input {
		if slice[i] == "+" || slice[i] == "-" || slice[i] == "*" {
			operators = append(operators, slice[i])
			alert = false
		} else {
			numbers = append(numbers, slice[i])
		}
	}
	if len(numbers) == 0 || len(operators) == 0 {
		return []int{}
	} else if len(operators) == 0 && len(numbers) == 1 {
		num, _ := strconv.Atoi(numbers[0])
		answer := []int{num}
		return answer
	}
	return RecurCount(numbers, operators, 0, len(numbers)-1)
}
