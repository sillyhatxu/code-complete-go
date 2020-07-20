package main

import (
	"fmt"
	"github.com/sillyhatxu/code-complete-go/design-patterns/behavioral-patterns/Observer/observer"
	"github.com/sillyhatxu/code-complete-go/design-patterns/behavioral-patterns/Observer/subject"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	trap := &observer.Observer{Name: "trap", ObserverChan: make(chan string), WG: wg}
	go trap.Observe()
	time.Sleep(20 * time.Microsecond)
	monster := &observer.Observer{Name: "monster", ObserverChan: make(chan string), WG: wg}
	go monster.Observe()
	time.Sleep(20 * time.Microsecond)
	treasure := &observer.Observer{Name: "treasure", ObserverChan: make(chan string), WG: wg}
	go treasure.Observe()
	hero := &subject.DefineSubject{MU: &sync.Mutex{}}
	hero.AttachObserver(trap.ObserverChan)
	hero.AttachObserver(monster.ObserverChan)
	hero.AttachObserver(treasure.ObserverChan)
	wg.Add(hero.GetObserverLength())
	hero.NotifyObservers("move")
	fmt.Println("notify observers waiting")
	wg.Wait()
	fmt.Println("notify observers end")
}
