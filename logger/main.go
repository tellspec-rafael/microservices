package main

import (
	"fmt"
	"log"

	zmq "github.com/pebbe/zmq4"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {

	log.SetOutput(&lumberjack.Logger{
		Filename:   "logs/log.log",
		MaxSize:    2, // megabytes
		MaxBackups: 10,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	})

	s, err := zmq.NewSocket(zmq.PULL)
	defer s.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	s.Bind("tcp://*:1")

	for {
		a, err := s.Recv(0)
		if err != nil {
			fmt.Println(err)
			return
		}
		log.Print(a)
	}
}
