// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b, c string = "b", "c"
	var d bool = true
	e := 100
	f := "f"
	j := true

	fmt.Printf("%T - %v\n", a, a)
	fmt.Printf("%T - %v\n", b, b)
	fmt.Printf("%T - %v\n", c, c)
	fmt.Printf("%T - %v\n", d, d)

	fmt.Printf("%T - %v\n", e, e)
	fmt.Printf("%T - %v\n", f, f)
	fmt.Printf("%T - %v\n", j, j)

}
func inc() {
	fmt.Println(e)
}
