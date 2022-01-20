package main

func swap_num(a int, b int) (int, int) {
	a, b = b, a
	return a, b
}

func main() {
	println(swap_num(11, 12))
}
