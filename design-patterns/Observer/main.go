package main

import (
	"github.com/sillyhatxu/code-complete-go/design-patterns/Observer/observer"
	"github.com/sillyhatxu/code-complete-go/design-patterns/Observer/subject"
	"time"
)

func main() {
	//initial object
	trap := &observer.Observer{Name: "trap", ObserverChan: make(chan string)}
	go trap.Observe()
	monster := &observer.Observer{Name: "monster", ObserverChan: make(chan string)}
	go monster.Observe()
	treasure := &observer.Observer{Name: "treasure", ObserverChan: make(chan string)}
	go treasure.Observe()
	hero := &subject.DefineSubject{}
	hero.AttachObserver(trap.ObserverChan)
	hero.AttachObserver(monster.ObserverChan)
	hero.AttachObserver(treasure.ObserverChan)

	hero.NotifyObservers("move")
	time.Sleep(5 * time.Second)
}
