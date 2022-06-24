package example

import "fmt"

var a int = 10

func Foo(n int) int {
	a := 1
	a += n
	return a
}

func PrintA() {
	fmt.Printf("%d\n", a)
}
