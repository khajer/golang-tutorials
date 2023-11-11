package main

import (
	"fmt"
	"log"
	"time"
)

func notiLog(word string) {
	fmt.Println(word, time.Now())
}

func main() {

	defer func() {
		fmt.Println()
		if r := recover(); r != nil { // recover must have in defer
			fmt.Println("recover", r)
		}
	}()
	fmt.Println("begin time", time.Now())
	defer notiLog("end")

	fmt.Println("hello")
	var inputStr string
	_, err := fmt.Scanf("%s", &inputStr)
	if err != nil {
		log.Fatal(err)
	}

	//defer
	// defer func() {
	// 	fmt.Println("test")
	// }()

	// pannic
	panic("test panic")

	//recovery

}
