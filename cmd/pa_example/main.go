package main

import (
	"fmt"
	"math/rand"
	"port/service/locklessq"
	"port/service/port"
	"sync/atomic"
	"time"
	"unsafe"
)

type noizeStreamImp struct {
	amp int32
	q   *locklessq.Q
}

//mono float 32
func (s *noizeStreamImp) Cbb(inputBuffer, outputBuffer unsafe.Pointer, frames uint64) {
	ob := (*[256]float32)(outputBuffer)
	//fmt.Println("frames ", frames)
	for i := uint64(0); i < frames; i++ {
		//f := (float32(atomic.LoadInt32(&s.amp)) / 1000)
		//fmt.Println("amplitude ", f)
		(*ob)[i] = s.q.Pop()
	}
}

func (s *noizeStreamImp) noizeGen() {
	for {
		if !s.q.Insert(rand.Float32()) {
			println("buffer full")
			time.Sleep(300 * time.Millisecond)
		}
	}

}

func (s *noizeStreamImp) volMod() {
	for {
		time.Sleep(1 * time.Millisecond)
		i := atomic.AddInt32(&s.amp, -1)
		if i == 0 {
			atomic.AddInt32(&s.amp, 1000)
		}
	}
}

func main() {
	fmt.Println("", pa.VersionText())
	pa.Initialize()
	pa.ListDevices()
	si := &noizeStreamImp{amp: 1000, q: locklessq.New(44100)}
	pa.Cba[0] = si
	//float 32
	s, err := pa.OpenDefaultStream(0, 1, pa.Float32, 44100, 256, nil)
	if s == nil {
		fmt.Println("sdsdsd", err)
		return
	}
	s.Start()
	go si.volMod()
	go si.noizeGen()
	time.Sleep(10 * time.Second)
	s.Stop()
	s.Close()
	return
}
