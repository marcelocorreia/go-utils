package main

import "fmt"

var VERSION string

func main() {
	if VERSION == "" {
		VERSION = "dev"
	}
	fmt.Println(VERSION)
}
