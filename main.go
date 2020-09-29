// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

const appName = "Golang Less"
const (
	_ = iota
	B = iota
	C = iota
	D = iota + 5
	E = iota
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	x, _, _ := TepesGo()
	fmt.Println(x)
}

func TepesGo() (int, bool, string) {
	return 100, false, "string"
}
