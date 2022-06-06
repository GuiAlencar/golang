package main

import "fmt"

func generica(inrterf interface{}) {
	fmt.Println(inrterf)
}

func main() {
	generica("string")
}
