package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var dbData = []int{1, 2, 3, 4, 5}
var wg = sync.WaitGroup{}

// var m = sync.Mutex{}
var m = sync.RWMutex{}
var results = []string{}

func main() {

	t0 := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Println("Time taken: ", time.Since(t0))

	fmt.Println("Results: ", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 1000

	time.Sleep(time.Duration(delay) * time.Millisecond)

	fmt.Println("DB Call ", i)

	save(fmt.Sprintf("DB Call %d", i))
	log()
	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Println("current Results: ", results)
	m.RUnlock()
}
