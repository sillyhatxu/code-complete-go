package singleton

import "sync"

var instantiated *single = nil

type single struct {
	Test string
}

var once sync.Once

func New() *single {
	once.Do(func() {
		instantiated = &single{}
	})
	return instantiated
}
