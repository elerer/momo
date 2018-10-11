package noizegen

import (
	"math"
	"testing"
)

func BenchmarkRead(b *testing.B) {
	var size int32 = 44100
	imp := New(size, 1024, 1024)

	//time.Sleep(2 * time.Second)

	buf := make([]byte, 1024)

	for n := 0; n < b.N; n++ {
		imp.Read(buf)
	}
	imp.Stop = true

	//time.Sleep(1 * time.Second)

	println("reads ", imp.Reads, " full ", imp.Full)
}

func BenchmarkFtoB(b *testing.B) {
	var ba [4]byte
	for n := 0; n < b.N; n++ {
		f := math.Float32bits(1.0)
		ba[0] = byte(f)
		ba[1] = byte(f >> 8)
		ba[2] = byte(f >> 16)
		ba[3] = byte(f >> 24)
	}

}
