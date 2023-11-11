package main

import (
	"fmt"

	tools "github.com/khajer/tools"
)

func main() {
	fmt.Println("unit test")
	data := tools.CheckData()
	fmt.Println(data)
}
