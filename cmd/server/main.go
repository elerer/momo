package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"port/service/noizegen"

	"github.com/gorilla/websocket"
)

//#include <string.h>
import "C"

var upgrader = websocket.Upgrader{} // use default options

func noize(w http.ResponseWriter, r *http.Request) {

	var buf_size int = 256 * 4
	fmt.Println("Starting noize osc")

	si := noizegen.New(44100, int32(buf_size), 1000)
	b := make([]byte, buf_size, buf_size)
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	i := 0
	for {
		si.Read(b)
		err = c.WriteMessage(websocket.BinaryMessage, b)
		println("wropte payload ", i)
		i++
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

var addr = flag.String("addr", "localhost:8088", "http service address")

func main() {
	flag.Parse()

	http.HandleFunc("/noize", noize)
	log.Fatal(http.ListenAndServe(*addr, nil))

	return
}
