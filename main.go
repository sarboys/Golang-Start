// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	a, b, c := 10, 20, 30
	if a > b {
		fmt.Println("a")
	} else if b > c {
		fmt.Println("b")
	} else {
		fmt.Println("c")
	}
	text := "admin"
	switch text {
	case "admin":
		fmt.Println("Cовпадает")
		fallthrough
	case "odmin":
		fmt.Println("Не совпадает2")
	case "odmi2n":
		fmt.Println("Не совпадает3")
	default:
		fmt.Println("SOME LOGIN")
	}

	for x := 1; x < 100; x++ {
		if x%10 == 0 {
			fmt.Println("Число", x)
		}
		fmt.Println(x)
	}
	x := 1
	for x > 0 {
		fmt.Println("x>0")
	}
}
