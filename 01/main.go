package main

import "fmt"

func main() {
	messageGreeting := "hello, please input your information"
	var finalMessage = "you information is "
	fmt.Println(messageGreeting)

	var name string
	var age int
	fmt.Print("input you name: ")
	// fmt.Scan(&name)
	//fmt.Scanf("%s", &name)
	fmt.Scanf("%s\n", &name)

	fmt.Print("input you age: ")
	fmt.Scanf("%d", &age)

	fmt.Printf("%s name: %s, age: %d ", finalMessage, name, age)

}
