package main

import "unsafe"

func main() {
	const const1 int = 12
	const const2 = "kkkk"

	println(len(const2),unsafe.Sizeof(const2))
}
