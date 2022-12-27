package main

import (
	"fmt"
	"github.com/likecodingloveproblems/pistachiowarehousemanagement"
)

func run() string {
	return pistachiowarehousemanagement.Shout("This is an example")
}

func main() {
	fmt.Println(run())
}