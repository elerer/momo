package main

import (
	"port/service/locklessq"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var size int32 = 44100
	var pushes int32 = 80000
	var pops int32 = 50000

	q := locklessq.New(size)
	var f float32 = 1.0
	start := false
	go func() {
		defer wg.Done()

		for i := 0; int32(i) < pushes; i++ {
			q.Insert(f)
			start = true
			f++
		}
		println("fff")
	}()

	go func() {
		defer wg.Done()

		for i := 0; int32(i) < pops; i++ {
			if start && q.Pop() == 0 {
				print("got 0 q has %d free space", q.ReadAvailble())
				return
			}
		}
		println("ssss")

	}()

	wg.Wait()

}
