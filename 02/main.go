package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	// var logined bool

	// var name string

	// var money int // 8/16/32/64

	// var age uint
	// var uIntPoint uintptr

	// var buffer byte

	// var datarune rune // alias for int32 // unitcode

	// var height float64

	// var datacomplex complex64 // complex 128

	// check type

	// func typeof(v interface{}) string {
	// 	switch v.(type) {
	// 	case int:
	// 		return "int"
	// 	case float64:
	// 		return "float64"
	// 	//... etc
	// 	default:
	// 		return "unknown"
	// 	}
	// }
	

	// define
	var Lat, Long float64 = 100.103, 13.001
	fmt.Println(Lat, Long)

	// convert
	const NAME = "kha"
	fmt.Println(NAME)

	var input_text string
	fmt.Println("input text")
	fmt.Scan(&input_text)
	fmt.Println("> " + input_text)
	fmt.Println(reflect.TypeOf(input_text))

	var age int
	fmt.Println("input you age")
	fmt.Scan(&input_text)
	var textAge = "you age is " + strconv.Itoa(3)
	fmt.Println(textAge)
	fmt.Println(reflect.TypeOf(age))
	x := strconv.Itoa(3)
	fmt.Println(x)

	// convert string to int
	intAge, err := strconv.Atoi("50")
	if err == nil {
		fmt.Println("error parse int")
	}
	fmt.Println(intAge)
	fmt.Println(strconv.Itoa(3))

	// convert string to float
	height := 1.30
	fmt.Println(height)
	txtH, err := strconv.ParseFloat("1.4232", 32)
	if err == nil {
		fmt.Println(txtH)
	}
	s := fmt.Sprintf("%f", 123.456) // %.2f, %8.2f
	fmt.Println(s)

	var src = 300
	var div = 300.0
	v := src / int(div)

	vf := float64(src) / div // src/div err. src > div
	fmt.Println(uint(300.0))
	fmt.Println(v, vf)


	var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n", 

}
