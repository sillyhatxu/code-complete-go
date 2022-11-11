package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"time"
)

func main() {
	fmt.Println("-------------------- Prepare Next TOTO Number --------------------")
	totoPool := make(map[int]bool)
	for i := 1; i <= 49; i++ {
		totoPool[i] = true
	}
	for {
		n, _ := rand.Int(rand.Reader, big.NewInt(49))
		delete(totoPool, int(n.Int64()))
		if len(totoPool) == 6 {
			break
		}
	}
	var nextLottery []int
	for selectedNumber := range totoPool {
		nextLottery = append(nextLottery, selectedNumber)
	}
	sort.Ints(nextLottery)
	fmt.Print(nextLottery)
	fmt.Println("\n", "-------------------- OK. Got it. I'm going to get the first prize. --------------------")
}

func backups() {
	fmt.Println("-------------------- Prepare Next TOTO Number --------------------")
	totoPool := make(map[int]bool)
	for {
		time.Sleep(time.Second)
		n, _ := rand.Int(rand.Reader, big.NewInt(49))
		totoPool[int(n.Int64()+1)] = true
		if len(totoPool) == 6 {
			break
		}
		fmt.Println("-------------------- got it --------------------")
	}
	var nextLottery []int
	for selectedNumber := range totoPool {
		nextLottery = append(nextLottery, selectedNumber)
	}
	sort.Ints(nextLottery)
	fmt.Print(nextLottery)
	fmt.Println("\n", "-------------------- OK. Got it. I'm going to get the first prize. --------------------")
}
