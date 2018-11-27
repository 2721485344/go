package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex
var val = 0
func main() {
	m = new(sync.RWMutex)
	go read(1)
	go write(2)
	go read(3)
	time.Sleep(time.Second * 2)
	}
func read(i int) {
	m.RLock()
	println("val: ", val)
	m.RUnlock()
}
func write(i int) {
	m.Lock()
	val = 10
	m.Unlock()
	}
