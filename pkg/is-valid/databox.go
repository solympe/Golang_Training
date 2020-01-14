package is_valid

type dataBox struct {
	data string
}

// IsValid ...
func (d *dataBox) IsValid() bool {

	if len(d.data) == 0 {
		return true
	}
	stack := []uint8{}

	stack = append(stack, d.data[0])
	j := 1
	for i := 1; i < len(d.data); i++ {
		if j == 0 || d.data[i]-stack[j-1] > 2 || d.data[i] == stack[j-1] {
			stack = append(stack, d.data[i])
			j++
		} else {
			stack = stack[:len(stack)-1]
			j--
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// NewDataBoxValidator ...
func NewDataBoxValidator(data string) DataBoxValidator {
	return &dataBox{data: data}
}
