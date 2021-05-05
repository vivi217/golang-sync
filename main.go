package main

import (
	"fmt"
	"sync"
	"time"
)

var oneDB sync.Mutex

func goFunc1() {
	oneDB.Lock()
	//fmt.Printf("begine  get data 1")
	time.Sleep(3)
	fmt.Printf("begine  get data 1\n")
	oneDB.Unlock()
}

func goFunc2() {
	oneDB.Lock()
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("begine  get data 2\n")
	oneDB.Unlock()

}

func goFunc3() {
	oneDB.Lock()
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Printf("begine  get data 3\n")
	oneDB.Unlock()
}

func main() {
	go goFunc1()
	go goFunc2()
	go goFunc3()
	time.Sleep(time.Duration(10) * time.Second)

	return
}
