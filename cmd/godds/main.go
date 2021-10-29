package main

import (
	"fmt"
	"time"

	"github.com/suutaku/godds/pkg/node"
)

func main() {
	c1 := node.NewConnector("", "", "")
	go func() {
		c2 := node.NewConnector("", "", "")
		c2.Listen()
		for {
			b, err := c2.Read()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(b))
		}
	}()
	c1.Dial()
	for {
		fmt.Println("send")
		c1.Write([]byte("hello,john"))
		time.Sleep(1 * time.Second)
	}
}
