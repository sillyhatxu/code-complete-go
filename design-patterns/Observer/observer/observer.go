package observer

import "fmt"

type Observer struct {
	Name         string
	ObserverChan chan string
}

func (obs *Observer) Observe() {
	event := <-obs.ObserverChan
	fmt.Printf("Observer Name : %s ; Event: %s\n", obs.Name, event)
}
