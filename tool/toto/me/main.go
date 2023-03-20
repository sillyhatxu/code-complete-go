package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("-------------------- Prepare Next TOTO Number --------------------")
	selected := 8
	totoPool := make([]int, 49, 49)
	//[25 31 20 44 29 22 43 1 42 24 23 37 8 45 47 15 16 30 26 6 33 36 49 4 11 13 28 12 14 7 40 38 34 9 32 48 35 5 21 39 17 10 19 46 18 27 2 41 3]
	//[25 31 20 44 29 22 43 1]
	//[32 27 16 13 25 21 39 40]
	//[10 5 9 34 31 46 3 7]
	for i := 0; i < 49; i++ {
		totoPool[i] = i + 1
	}
	rand.Seed(time.Now().UnixNano())
	for end := len(totoPool) - 1; end >= 0; end-- {
		randNum := rand.Intn(end + 1)
		totoPool[end], totoPool[randNum] = totoPool[randNum], totoPool[end]
	}
	fmt.Println(totoPool)
	fmt.Println(totoPool[0:selected])
	fmt.Println("\n", "-------------------- OK. Got it. I'm going to get the first prize. --------------------")
}

//func RandomVersion2() {
//	fmt.Println("-------------------- Prepare Next TOTO Number --------------------")
//	totoPool := make(map[int]bool)
//	for i := 1; i <= 49; i++ {
//		totoPool[i] = true
//	}
//	for {
//		n, _ := rand.Int(rand.Reader, big.NewInt(49))
//		delete(totoPool, int(n.Int64()))
//		if len(totoPool) == 6 {
//			break
//		}
//	}
//	var nextLottery []int
//	for selectedNumber := range totoPool {
//		nextLottery = append(nextLottery, selectedNumber)
//	}
//	sort.Ints(nextLottery)
//	fmt.Print(nextLottery)
//	fmt.Println("\n", "-------------------- OK. Got it. I'm going to get the first prize. --------------------")
//}

//func RandomVersion1() {
//	fmt.Println("-------------------- Prepare Next TOTO Number --------------------")
//	totoPool := make(map[int]bool)
//	for {
//		time.Sleep(time.Second)
//		n, _ := rand.Int(rand.Reader, big.NewInt(49))
//		totoPool[int(n.Int64()+1)] = true
//		if len(totoPool) == 6 {
//			break
//		}
//		fmt.Println("-------------------- got it --------------------")
//	}
//	var nextLottery []int
//	for selectedNumber := range totoPool {
//		nextLottery = append(nextLottery, selectedNumber)
//	}
//	sort.Ints(nextLottery)
//	fmt.Print(nextLottery)
//	fmt.Println("\n", "-------------------- OK. Got it. I'm going to get the first prize. --------------------")
//}
