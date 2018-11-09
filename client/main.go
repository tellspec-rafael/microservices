//
// Hello World Zeromq Client
//
// Author: Aaron Raddon   github.com/araddon
// Requires: http://github.com/alecthomas/gozmq
//
package main

import (
	"fmt"
	"sync"

	zmq "github.com/pebbe/zmq4"
)

func clientFunc(wg *sync.WaitGroup, index int) {
	socket, _ := zmq.NewSocket(zmq.REQ)
	defer socket.Close()

	fmt.Printf("Client %d created\n", index)
	socket.Connect("ipc:///router/router.ipc")
	//socket.Connect("tcp://127.0.0.1:5559")

	for i := 0; i < 10; i++ {
		// send hello
		msg := fmt.Sprintf("Hello %d", i)
		socket.Send(msg, 0)
		fmt.Printf("Client %d Sending %v\n", index, msg)

		// Wait for reply:
		reply, _ := socket.Recv(0)
		fmt.Printf("Client %d Received %v\n", index, string(reply))
	}
	defer wg.Done()
}

func main() {
	numberOfClients := 16
	var wg sync.WaitGroup
	wg.Add(numberOfClients)
	for index := 0; index < numberOfClients; index++ {
		go clientFunc(&wg, index)
	}
	wg.Wait()
}
