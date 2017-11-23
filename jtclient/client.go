// Package jtclient provides a client for making commands to jt server.
package jtclient

import (
	"net"
)

// Client is struct that contains tcp connection to jt server
// and some configuration info.
type Client struct {
	options Options
	conn    net.Conn
}

type Options struct {
	Host, Password string
}

// NewClient creates new jt client
func NewClient(o Options) (*Client, error) {
	conn, err := net.Dial("tcp", o.Host)
	if err != nil {
		return &Client{}, err
	}

	c := &Client{
		options: o,
		conn:    conn,
	}
	return c, nil
}
