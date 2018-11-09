package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
)

func main() {

	fmt.Println("Router started.")

	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer frontend.Close()
	defer backend.Close()
	frontend.Bind("ipc://router.ipc")
	backend.Bind("tcp://*:1")

	// Start router see http://api.zeromq.org/4-1:zmq-proxy#toc2
	err := zmq.Proxy(frontend, backend, nil)
	log.Fatalln("Proxy interrupted:", err)
}
