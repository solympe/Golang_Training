package singleton

import "sync"

var (
	once      sync.Once
	singleMan *singleton
)

// SyngletonMaker ...
type SingleMaker interface {
	SetName(name string)
	SetSurname(name string)
}

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
func Singleton() *singleton {
	once.Do(func() {
		singleMan = &singleton{}
	})
	return singleMan
}
