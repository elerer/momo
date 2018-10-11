package utils

import "math"

func Float32ToBytes(fl float32, b []byte, j int) {
	f := math.Float32bits(fl)
	b[j] = byte(f)
	b[j+1] = byte(f >> 8)
	b[j+2] = byte(f >> 16)
	b[j+3] = byte(f >> 24)
}

func BytesToFloat32(b []byte, j int) float32 {
	var f uint32
	f |= uint32(b[j+3])
	f = f << 8
	f |= uint32(b[j+2])
	f = f << 8
	f |= uint32(b[j+1])
	f = f << 8
	f |= uint32(b[j])
	return math.Float32frombits(f)
}
