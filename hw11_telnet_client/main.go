package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	argsLen := len(os.Args)
	if argsLen < 3 || argsLen > 4 {
		log.Fatalf("wrong argument amount")
	}

	var timeout time.Duration
	flag.DurationVar(&timeout, "timeout", time.Second*10, "timeout for connection")
	flag.Parse()

	host := ""
	port := 0
	var err error

	if argsLen == 3 {
		host = os.Args[1]
		port, err = strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalln("valid port reqired")
		}
	} else {
		host = os.Args[2]
		port, err = strconv.Atoi(os.Args[3])
		if err != nil {
			log.Fatalln("valid port reqired")
		}
	}

	client := NewTelnetClient(fmt.Sprintf("%s:%d", host, port), timeout, os.Stdin, os.Stdout)

	if err := client.Connect(); err != nil {
		log.Fatalln("connection failed:", err)
	}
	defer client.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer cancel()

	go func() {
		err := client.Send()
		if err != nil {
			log.Fatalln("error sending")
		}
		cancel()
	}()

	go func() {
		err := client.Receive()
		if err != nil {
			log.Fatalln("error receiving")
		}
		cancel()
	}()

	<-ctx.Done()
}
