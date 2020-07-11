package observer

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

type Observer struct {
	Name         string
	ObserverChan chan string
	WG           *sync.WaitGroup
}

func (obs *Observer) Observe() {
	random, _ := rand.Int(rand.Reader, big.NewInt(5))
	sleep := random.Int64() + 1
	event := <-obs.ObserverChan
	time.Sleep(time.Duration(sleep) * time.Second)
	fmt.Printf("Observer Name : %s; Event: %s; e time : %d\n", obs.Name, event, sleep)
	obs.WG.Done()
}
