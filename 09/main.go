package main

import (
	"fmt"
	"sync"
	"time"
)

func Run1() {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, "0")
		}
		time.Sleep(time.Microsecond * 50)

	}
}
func Run2() {
	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Println(i, "1")
		}
		time.Sleep(time.Microsecond * 50)

	}

}

func Run3(c chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	c <- 421

}

func Run4(c chan string) {

	c <- "5r5555"
}

func RunWait1() {
	fmt.Println("wait 1")
	time.Sleep(time.Second * 2)

}

func RunWait2() {
	fmt.Println("wait 2")
	time.Sleep(time.Second * 2)

}

func main() {
	fmt.Println("run goroutine")
	go Run1()
	Run2()

	// time.Sleep(6 * time.Second)

	// val := go Run3() <-- cannot do this
	c := make(chan int)
	ch := make(chan string)
	go Run3(c)
	go Run4(ch)

	val, txt := <-c, <-ch
	println("xxxx")
	println(val, txt)

	// waiting all
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		RunWait1()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		RunWait2()
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("end")

}
