package subject

import (
	"sync"
)

type DefineSubject struct {
	observerArray []chan string
	mu            *sync.Mutex
}

func (ds *DefineSubject) validate() bool {
	if ds == nil || ds.observerArray == nil {
		return false
	}
	return true
}

func (ds *DefineSubject) AttachObserver(observer chan string) {
	ds.observerArray = append(ds.observerArray, observer)
}

func (ds *DefineSubject) DetachObserver(observer chan string) {
	if !ds.validate() {
		return
	}
	for i, v := range ds.observerArray {
		if v == observer {
			ds.observerArray = append(ds.observerArray[:i], ds.observerArray[i+1:]...)
			break
		}
	}
}

func (ds *DefineSubject) NotifyObservers(name string) {
	if !ds.validate() {
		return
	}
	for _, v := range ds.observerArray {
		v <- name
	}
}
