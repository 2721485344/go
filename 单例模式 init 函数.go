package main
import ( "fmt" )
var m *Manager
func GetInstance() *Manager {
	if m == nil {
		m = &Manager{}
	}
	return m
}
type Manager struct{}
func (p Manager) Manage() {
	fmt.Println("manage...")
}
func init() {
	fmt.Println("Hello INIT")
}
func main() {
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())
	fmt.Printf("%p\n", GetInstance())

	//-------------------------init 函数---------------
	fmt.Println("main")
}
