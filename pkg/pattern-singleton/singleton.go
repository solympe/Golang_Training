package singleton

import "sync"

var (
	once      sync.Once
	singleMan Singletoner
)

type singleton struct {
	name    string
	surname string
}

// SetName ...
func (s *singleton) SetName(name string) {
	s.name = name
}

// SetSurname ...
func (s *singleton) SetSurname(sname string) {
	s.surname = sname
}

// Singleton ...
func Singleton() Singletoner {
	once.Do(func() {
		singleMan = NewSingletoner()
	})
	return singleMan
}

// NewSingletoner ...
func NewSingletoner() Singletoner {
	return &singleton{}
}
