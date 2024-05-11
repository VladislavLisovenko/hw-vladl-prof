package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type client struct {
	address    string
	timeout    time.Duration
	in         io.ReadCloser
	out        io.Writer
	connection net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *client) Connect() error {
	connection, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return fmt.Errorf("connection error: %w", err)
	}
	fmt.Printf("connected to %s\n", c.address)
	c.connection = connection

	return nil
}

func (c *client) Send() error {
	if _, err := io.Copy(c.connection, c.in); err != nil {
		return fmt.Errorf("send error: %w", err)
	}
	fmt.Println("data sent")
	return nil
}

func (c *client) Receive() error {
	if _, err := io.Copy(c.out, c.connection); err != nil {
		return fmt.Errorf("receive error: %w", err)
	}
	fmt.Println("data were received")
	return nil
}

func (c *client) Close() error {
	if c.connection != nil {
		if err := c.connection.Close(); err != nil {
			return fmt.Errorf("connection closing error: %w", err)
		}
	}
	fmt.Println("connection closed")
	return nil
}
