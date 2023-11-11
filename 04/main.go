package main

import "fmt"

func load() {

}
func checkYouWinScore(score int) bool {
	if score > 0 && score < 10 {
		return true
	}
	return false
}

func returnText() string {
	return "hello"

}

func setText(txt *string) error {
	*txt = "hello"
	return nil
}

func last(fn func()) {
	fn()
}

// variadic function
func sum(num ...int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func MultiReturn() (int, int, int) {
	return 0, 0, 0
}
func main() {

	fmt.Println("function")
	load()

	fmt.Println(checkYouWinScore(20))

	// anonymous function => like closure
	func() {
		fmt.Println("test")
	}()

	last(func() {
		fmt.Println("test parameter function")
	})

	// pointer
	text := ""
	setText(&text)
	val := sum(1) + sum(1, 2)
	fmt.Println(val)

	// closure => anonymous funciton

}
