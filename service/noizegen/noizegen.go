package noizegen

import (
	"math/rand"
	"port/service/locklessq"
	"port/service/utils"
	"time"
)

//#include <string.h>
import "C"

type Imp struct {
	Amp    int32
	Q      *locklessq.Q
	buffer []float32
	Full   int
	Reads  int
	Stop   bool
}

func New(qSize, bufsize, amp int32) *Imp {
	imp := &Imp{Amp: amp, Q: locklessq.New(qSize), buffer: make([]float32, qSize, qSize), Stop: false}
	go imp.noizeGen()
	go func() {
		for {
			dd := imp.Q.WriteAvailble()

			if dd > qSize-10 {
				println("free ", dd)
			}

		}

	}()
	time.Sleep(3 * time.Second)
	return imp
}

func (s *Imp) Read(b []byte) {
	if len(b)%4 != 0 {
		panic("argument should be devideable by 4")
	}
	size := len(b) / 4

	for i := 0; i < size; i++ {
		f := s.Q.Pop()
		if f == 0 {
			panic("--sending 0")
		}
		utils.Float32ToBytes(f, b, i*4)
	}
	s.Reads++
	//err := binary.Write(b, binary.LittleEndian, s.buffer[0])
	// C.memcpy(unsafe.Pointer(&b[0]), unsafe.Pointer(&s.buffer[0]), C.size_t(size))

}

func (s *Imp) noizeGen() {
	for {
		if !s.Q.Insert(rand.Float32()) {
			//println("buffer full")
			s.Full++
			time.Sleep(3 * time.Millisecond)
		}
		if s.Stop {
			return
		}
	}

}
