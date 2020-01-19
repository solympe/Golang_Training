package singleton

import "sync"

var (
	once      sync.Once
	singleMan Singleton
)

// Singleton ...
type Singleton interface {
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

// MakeSingleton ...
func MakeSingleton() Singleton {
	once.Do(func() {
		singleMan = NewSingletone()
	})
	return singleMan
}

// NewSingletone ...
func NewSingletone() Singleton {
	return &singleton{}
}
