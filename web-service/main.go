//+build prod, dev, test
package main

import (
	"fmt"
	"net/http"
  	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("heyo, world.")
	test()
}

// test function
func test() {
	fmt.Println("test function")
}