package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("华氏度f=%g and 摄氏度c=%g", f, c)
}
