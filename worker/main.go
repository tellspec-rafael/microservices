package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	zmq "github.com/pebbe/zmq4"
)

func worker(workerInternalIdentifier string) {
	// Socket to talk to dispatcher
	receiver, _ := zmq.NewSocket(zmq.REP)

	// REP socket monitor, all events
	// err := receiver.Monitor("tcp://logger:2.req", zmq.EVENT_ALL)
	// if err != nil {
	// 	fmt.Print("rep.Monitor:", err)
	// }

	defer receiver.Close()
	receiver.Connect("ipc:///tmp/workers" + workerInternalIdentifier + ".ipc")

	for true {
		received, _ := receiver.Recv(0)
		logger(fmt.Sprintf("Received request [%s]\n", received))

		// Do some 'work'
		time.Sleep(time.Second)

		// Send reply back to client
		receiver.Send(received, 0)
	}
}

func randomString() string {
	source := "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	target := make([]byte, 20)
	for i := 0; i < 20; i++ {
		target[i] = source[rand.Intn(len(source))]
	}
	return string(target)
}

func main() {
	threadsPtr := flag.Int("threads", 1, "Number of Threads")
	flag.Parse()

	// Create a random indentifier for internal address
	rand.Seed(time.Now().UTC().UnixNano())
	workerInternalIdentifier := randomString()

	// Start threads
	for i := 0; i != *threadsPtr; i = i + 1 {
		go worker(workerInternalIdentifier)
		logger(fmt.Sprintf("Thread %d created.\n", i))
	}

	frontend, _ := zmq.NewSocket(zmq.ROUTER)
	backend, _ := zmq.NewSocket(zmq.DEALER)
	defer frontend.Close()
	defer backend.Close()
	frontend.Connect("tcp://router:5558")
	backend.Bind("ipc:///tmp/workers" + workerInternalIdentifier + ".ipc")

	// Start worker router see http://api.zeromq.org/4-1:zmq-proxy#toc2
	err := zmq.Proxy(frontend, backend, nil)
	logger(fmt.Sprintf("Proxy interrupted: %s", err))
}

func logger(log string) {

	fmt.Printf(log)

	socket, _ := zmq.NewSocket(zmq.PUSH)
	defer socket.Close()
	socket.Connect("tcp://logger:1")
	socket.Send(log, 0)
	socket.Close()
}
