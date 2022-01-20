package main

func main() {
	//const const1 int = 12
	//const const2 = "kkkk"
	//
	//println(len(const2),unsafe.Sizeof(const2))

	//iota
	//const (
	//	a = iota
	//	b
	//	c
	//)
	//println(a,b,c)

	const (
		i = 2 << iota
		j
		k
		l
	)
	println(i, j, k, l)
}
