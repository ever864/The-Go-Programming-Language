package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//var s, sep string
	//for i := 1; i < len(os.Args); i++ {
	//	s += sep + os.Args[i]
	//	sep = " "
	//}
	//

	//s, sep := "", ""
	//for _, arg := range os.Args[1:] {
	//	s += sep + arg
	//	sep = " ever "
	//}

	fmt.Println(strings.Join(os.Args[1:], " "))
}
