package data

import (
	"fmt"
	"strconv"
)

// Data ...
type Data interface {
	Compute(input string)
	GetAnswer()
}

type data struct {
	answer []int
}

// GetAnswer ...
func (d *data) GetAnswer() {
	fmt.Println(d.answer)
}

// Compute ...
func (d *data) Compute(input string) {
	slice := d.splitInput(input)
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
		d.answer = answer
	}
	d.answer = d.recurCount(numbers, operators, 0, len(numbers)-1)
}

func (d *data) calculate(na int, x string, nb int) int {
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

func (d *data) recurCount(numbers []string, operator []string, iStart int, iEnd int) []int {
	result := []int{}

	if iStart == iEnd {
		intVal, _ := strconv.Atoi(numbers[iStart])
		result = append(result, intVal)
	} else {
		for iNow := iStart; iNow < iEnd; iNow++ {

			leftSide := d.recurCount(numbers, operator, iStart, iNow)
			rightSide := d.recurCount(numbers, operator, iNow+1, iEnd)

			for _, left := range leftSide {
				for _, right := range rightSide {
					result = append(result, d.calculate(left, operator[iNow], right))
				}
			}
		}
	}
	return result
}

func (d *data) splitInput(input string) (output []string) {
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

// NewData ...
func NewData() Data {
	return &data{
	}
}
