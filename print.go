package golangpractice
package main

import "fmt"

func printBody() {
	var name = "Go"
	language, version := "Go", 1.21
	fmt.Println(name, language, version)
}

func main() {
	printBody()
}
