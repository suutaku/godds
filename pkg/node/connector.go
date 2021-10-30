package node

import (
	"fmt"
	"net"
)

var (
	defaultAddress  = "224.0.0.1"
	defaultPort     = "56789"
	defaultProtocol = "udp"
)

type Connector struct {
	port     string
	proto    string
	addr     string
	dialer   net.Conn
	listener *net.UDPConn
	income   chan []byte
	running  bool
}

func NewConnector(addr string, port string, proto string) *Connector {
	ret := &Connector{
		income:  make(chan []byte),
		running: true,
	}
	if addr == "" {
		ret.addr = defaultAddress
	}
	if port == "" {
		ret.port = defaultPort
	}
	if proto == "" {
		ret.proto = defaultProtocol
	}
	return ret
}

func (c *Connector) Dial() {
	conn, err := net.Dial(c.proto, c.addr+":"+c.port)
	if err != nil {
		panic(err)
	}
	c.dialer = conn
	fmt.Println("Connector Dial")
}

func (c *Connector) Start() {
	c.Dial()
	c.Listen()
}

func (c *Connector) Listen() {
	udp_addr, err := net.ResolveUDPAddr("udp", c.addr+":"+c.port)
	if err != nil {
		panic(err)
	}
	conn, err := net.ListenMulticastUDP(c.proto, nil, udp_addr)
	if err != nil {
		panic(err)
	}
	c.listener = conn
	go func() {
		for c.running {
			buf := make([]byte, 1024)
			n, _, err := c.listener.ReadFrom(buf)
			if err != nil {
				fmt.Println(err)
			}
			if n > 0 {
				c.income <- buf[:n]
			}
		}
	}()
	fmt.Println("Connector listen")
}

func (c *Connector) Stop() {
	c.dialer.Close()
	c.listener.Close()
	c.running = false
}

func (c *Connector) Write(data []byte) error {
	_, err := c.dialer.Write(data)
	return err
}

func (c *Connector) Read() ([]byte, error) {
	buf := <-c.income
	return buf, nil
}
