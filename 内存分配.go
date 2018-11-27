package main

import (
	"bytes"
	"fmt"
	"sync"
)

type SyncedBuffer struct {
	lock sync.Mutex
	bbuffer bytes.Buffer
}
func main() {
	_ = new(SyncedBuffer)
	var _ SyncedBuffer
	var p1 *[]int = new([]int)
	var p2 []int = make([]int,10000)
	fmt.Printf("%p", p1)
	fmt.Printf("%p", p2)
	}
