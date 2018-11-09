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
	"time"

	zmq "github.com/pebbe/zmq4"
)

func clientFunc(wg *sync.WaitGroup, index int) {
	socket, _ := zmq.NewSocket(zmq.REQ)
	defer socket.Close()

	//fmt.Printf("Client %d created\n", index)
	fmt.Print("Created\n")
	socket.Connect("ipc:///router/router.ipc")
	//socket.Connect("tcp://127.0.0.1:5559")

	for i := 0; i < 10; i++ {
		// send hello
		msg := fmt.Sprintf("Hello %d", i)
		socket.Send(msg, 0)
		//fmt.Printf("Client %d Sending [%v]\n", index, msg)
		fmt.Printf("Sending [%v]\n", msg)

		// Wait for reply:
		reply, _ := socket.Recv(0)
		//fmt.Printf("Client %d Received [%v]\n", index, string(reply))
		fmt.Printf("Received [%v]\n", string(reply))
	}
	defer wg.Done()
}

func main() {
	numberOfClients := 1
	var wg sync.WaitGroup
	wg.Add(numberOfClients)
	start := time.Now()
	for index := 0; index < numberOfClients; index++ {
		go clientFunc(&wg, index)
	}
	wg.Wait()
	fmt.Printf("It took: %f seconds.\n", time.Now().Sub(start).Seconds())
}
