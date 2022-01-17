package main

import (
	"unicode/utf8"
)

func main(){
	println(len("你好")) // 6
	println(utf8.RuneCountInString("你好")) // 2
	println(len("abcd")) // 4

}

