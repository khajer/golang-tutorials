package main

import (
	"time"

	pgb "github.com/khajer/minpgb"
)

func main() {

	pb := pgb.New()
	pb.Total = 100

	for i := 0; i < 100; i++ {
		curr := pb.GetCurrent()
		pb.SetCurrent(curr + 1)
		time.Sleep(50 * time.Millisecond)
	}
	pb.End()
}
