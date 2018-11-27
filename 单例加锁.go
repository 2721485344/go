package main

import (
	"fmt"
	"sync"
	"time"
)

var m *Manager
var lock *sync.Mutex = &sync.Mutex{}
func GetInstance() *Manager {
	lock.Lock()
	defer lock.Unlock()
	if m == nil {
		m = &Manager{}
		}
	return m
}
type Manager struct{}
func (p Manager) Manage() {
	fmt.Println("manage...")
}
func main() {
	go func() {fmt.Printf("%p\n", GetInstance())}()
	go func() { fmt.Printf("%p\n", GetInstance()) }()
	go func() { fmt.Printf("%p\n", GetInstance()) }()
	go func() { fmt.Printf("%p\n", GetInstance()) }()
	go func() { fmt.Printf("%p\n", GetInstance()) }()
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())
	time.Sleep(time.Second)
}
