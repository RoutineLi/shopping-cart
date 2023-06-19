package main

import (
	"fmt"
	"graduate_design/pkg"
)

func main() {
	key := "c0bb490b-0a81-4f7a-9008-9bcb93fa90f0"
	sec := "4e37b29c-1494-4422-b907-167ad043ed45d"
	fmt.Println(pkg.Md5(key + sec))

}
