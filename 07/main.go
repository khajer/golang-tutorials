package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")

	// array
	primes := [6]int{2, 3, 5, 7, 11, 13} // create array with init value
	fmt.Println(primes)

	s := make([]string, 3) // create array
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	ss := make([]int, 0, 3) // create array
	for i := 0; i < 30; i++ {
		ss = append(ss, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(ss), len(ss), ss)
	}

	// slide
	fmt.Println(ss)
	fmt.Println(ss[0:10]) // start:end
	fmt.Println(ss[:10])
	fmt.Println(ss[5:10])
	fmt.Println(ss[10:len(ss)])
	fmt.Println(ss[10:])
	fmt.Println(ss[:])

	txt := "hello, world"
	txtHello := txt[:5]
	txtWorld := txt[7:]

	fmt.Println(txtHello, txtWorld)

}
