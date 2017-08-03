package main

import (
	"fmt"
	"sync"
)

var num int = 0

func increase(lock *sync.Mutex) {
	lock.Lock()
	num = num + 1
	fmt.Println(num)
	lock.Unlock()
}

func main() {
	lock := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go increase(lock)
	}

	for {
		lock.Lock()
		c := num
		lock.Unlock()
		if c == 10 {
			break
		}
	}
}
