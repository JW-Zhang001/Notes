package example

import "fmt"

var a int = 10

// Foo 值传递,传递进来的参数只会影响局部数据,不影响全局数据
func Foo(n int) int {
	a := 1
	a += n
	return a
}

func PrintA() {
	fmt.Printf("%d\n", a)
}
