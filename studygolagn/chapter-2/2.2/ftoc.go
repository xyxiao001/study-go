package main

import "fmt"

func main() {
	fmt.Printf("华氏度f=%g and 摄氏度c=%g\n", 212.0, fToC(212))
	fmt.Printf("华氏度f=%g and 摄氏度c=%g\n", 32.0, fToC(32.0))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
