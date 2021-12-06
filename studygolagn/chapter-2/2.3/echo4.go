// go run studygolagn/chapter-2/2.3/echo4.go -s / a bc def
// go run studygolagn/chapter-2/2.3/echo4.go -n a bc def
package main

import (
	"flag"
	"fmt"
	"strings"
)

// 注册命令行提示 -h 显示提示
var n = flag.Bool("n", false, "省略后换行符")
var sep = flag.String("s", " ", "分隔符")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
