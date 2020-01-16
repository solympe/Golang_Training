package data

import "strconv"

// Data ...
type Data interface {
	Compute() []int
}

type data struct {
	input string
}

// Compute ...
func (d *data) Compute() []int {
	var numbers []string
	var operators []string
	parsedInput := d.splitInput(d.input)

	for _, i := range parsedInput {
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
	return d.recurCount(numbers, operators, 0, len(numbers)-1)
}

func (d *data) calculate(num1 int, x string, num2 int) int {
	switch x {
	case "+":
		return num1 + num2
	case "-":
		return num1 - num2
	case "*":
		return num1 * num2
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
	var tmp string
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
func NewData(input string) Data {
	return &data{
		input: input,
	}
}
