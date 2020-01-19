package data

// Data ...
type Data interface {
	IsValid() bool
}

type data struct {
	scopes string
}

// IsValid ...
func (d *data) IsValid() bool {

	if len(d.scopes) == 0 {
		return true
	}
	stack := []uint8{}

	stack = append(stack, d.scopes[0])
	j := 1
	for i := 1; i < len(d.scopes); i++ {
		if j == 0 || d.scopes[i]-stack[j-1] > 2 || d.scopes[i] == stack[j-1] {
			stack = append(stack, d.scopes[i])
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

// NewData ...
func NewData(scopes string) Data {
	return &data{scopes: scopes}
}
