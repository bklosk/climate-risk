package main

import (
	"fmt"
	"net/http"
  	"github.com/gin-gonic/gin"
)

// +build prod, dev, test
func main() {
	fmt.Println("heyo, world.")
	test()
}

// test function
func test() {
	fmt.Println("test function")
}