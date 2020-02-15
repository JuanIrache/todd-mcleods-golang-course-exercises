package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("OS\t\t%v\n", runtime.GOOS)
	fmt.Printf("Architecture\t%v\n", runtime.GOARCH)
	fmt.Printf("Cores\t\t%d\n", runtime.NumCPU())
	fmt.Printf("Go version\t%v\n", runtime.Version())
}
