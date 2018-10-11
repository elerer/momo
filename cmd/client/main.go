package main

import (
	"fmt"
	"log"
	"port/service/locklessq"
	"port/service/port"
	"port/service/utils"
	"unsafe"

	"github.com/gorilla/websocket"
)

//#include <string.h>
import "C"

type streamImp struct {
	q   *locklessq.Q
	num int
}

//mono float 32
func (s *streamImp) Cbb(inputBuffer, outputBuffer unsafe.Pointer, frames uint64) {
	ob := (*[512]float32)(outputBuffer)
	//fmt.Println("frames ", frames)
	for i := uint64(0); i < frames; i++ {
		vaal := s.q.Pop()
		if vaal == 0 {
			println("--got 0")
		}
		//fmt.Println(s.num, " ", vaal)
		s.num++
		(*ob)[i] = vaal
	}
	//panic("imaaaaaaa")
}

func main() {
	//fmt.Println("", pa.VersionText())

	si := &streamImp{q: locklessq.New(88200)}

	pa.Cba[0] = si

	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8088/noize", nil)
	if err != nil {
		log.Fatal("Dial Error: ", err)
	}
	fmt.Println("Connected to ws server")
	pa.Initialize()
	pa.ListDevices()
	s, _ := pa.OpenDefaultStream(0, 1, pa.Float32, 44100, 512, nil)
	s.Start()
	defer func() {
		conn.Close()
		s.Stop()
		s.Close()
	}()

	fmt.Println("Stream is open")

	for {
		_, bytes, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket closed.")
			return
		}
		recBytes := len(bytes)
		if si.q.WriteAvailble() >= int32(recBytes) {
			var f float32
			for i := 0; i < recBytes/4; i++ {
				// C.memcpy(unsafe.Pointer(&f), unsafe.Pointer(&bytes[i*4]), C.size_t(4))
				if f = utils.BytesToFloat32(bytes, i*4); f == 0 {
					println("--go 0")
				}
				si.q.Insert(f)
			}

		} else {
			println("q is full, tienes ", si.q.ReadAvailble(), " para read")
		}
	}

	return
}
