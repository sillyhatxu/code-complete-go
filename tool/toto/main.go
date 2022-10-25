package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("-------------------- Current TOTO Number --------------------")
	randNumbers := rand.Perm(49)
	start := rand.Intn(44)
	var nextLottery []int
	for _, randNumber := range randNumbers[start : start+6] {
		nextLottery = append(nextLottery, randNumber+1)
	}
	sort.Ints(nextLottery)
	fmt.Print(nextLottery)
	fmt.Println("\n", "-------------------- Current TOTO Number --------------------")
}
