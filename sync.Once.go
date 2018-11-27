package main

import (
	"fmt"
	"sync"
	"time"
)

var m *Manager
var once sync.Once
func GetInstance() *Manager {
	once.Do(func() {
		m = &Manager{} })
	return m
}
type Manager struct{}
func (p Manager) Manage() {
	fmt.Println("manage...")
}
func main() {
	go func() { fmt.Printf("%p\n", GetInstance()) }()
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
