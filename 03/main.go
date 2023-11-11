package main

import (
	"fmt"
	"time"
)

func main() {
	// -----------------------------------
	// loop
	// -----------------------------------
	for i := 0; i < 20; i++ {
		fmt.Println(i)
	}
	/*
		for {

		}
	*/

	numbers := []int{1, 2, 3, 4, 5}
	for idx, value := range numbers {
		fmt.Println(idx, value)
	}

	maxCount := 5

	for maxCount < 100 {
		maxCount++

		if maxCount > 10 {
			break
		}
		fmt.Println(maxCount)
	}

	// -----------------------------------
	// if
	// -----------------------------------
	score := 20

	if score < 10 {
		fmt.Println("score < 10 ")
	} else if score > 10 {
		fmt.Println("score > 10")
	} else {
		fmt.Println("score don't known")
	}

	// switch case with condition
	switch score {
	case 10, 50:
		{
			fmt.Println("switch_10")
		}
	case 20:
		{
			fmt.Println("switch_20")
		}
	default:
		{
			fmt.Println("switch_default")
		}
	}

	// switch with no condition
	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
	case hour < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

}
