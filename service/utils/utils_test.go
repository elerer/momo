package utils

import "testing"

func TestToFromBytes(t *testing.T) {
	var f float32 = 1839283.345
	b := make([]byte, 4)
	Float32ToBytes(f, b, 0)
	f2 := BytesToFloat32(b, 0)

	if f != f2 {
		t.Errorf("expected equal %f and %f", f, f2)
	}
}
