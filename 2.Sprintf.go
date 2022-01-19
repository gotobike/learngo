package main

import "fmt"

/*// Sprintf formats according to a format specifier and returns the resulting string.
func Sprintf(format string, a ...interface{}) string {
	p := newPrinter()
	p.doPrintf(format, a)
	s := string(p.buf)
	p.free()
	return s
}*/

func main() {
	var stockcode = 123
	var enddate = "2020-12-31"
	var target_url = fmt.Sprintf("Code = %d&endDate = %s", stockcode, enddate)
	fmt.Println(target_url)
}
