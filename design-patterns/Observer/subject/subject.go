package subject

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/Observer/observer"
	"sync"
)

type DefineObservable struct {
	observerArray []observer.DefineObserver
	mu            *sync.Mutex
}

func (do *DefineObservable) validate() bool {
	if do == nil || do.observerArray == nil {
		return false
	}
	return true
}

func (do *DefineObservable) attachObserver(o observer.DefineObserver) {
	do.observerArray = append(do.observerArray, o)
}

func (do *DefineObservable) detachObserver(o observer.DefineObserver) {
	if !do.validate() {
		return
	}
	for i, v := range do.observerArray {
		if v == o {
			do.observerArray = append(do.observerArray[:i], do.observerArray[i+1:]...)
		}
	}
}

func (do *DefineObservable) notifyObservers() {
	if !do.validate() {
		return
	}
	for _, v := range do.observerArray {
		v.update()
	}
}
