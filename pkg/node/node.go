package node

import (
	"context"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Node struct {
	publisher  *Publisher
	subscriber *Subscriber
	running    bool
}

func NewNode(ctx context.Context, addr string, port string) *Node {
	conn := NewConnector(addr, port, "")
	pb := NewPublisher(conn)
	sb := NewSubscriber(ctx, conn)
	return &Node{
		publisher:  pb,
		subscriber: sb,
	}
}

func (n *Node) Start() {
	n.publisher.Start()
	go n.subscriber.Start()
	n.running = true
}

func (n *Node) Stop() {
	n.subscriber.Stop()
	n.running = false
}

func (n *Node) IsRunning() bool {
	return n.running
}

func (n *Node) Subscribe(msg proto.Message, callback func(*anypb.Any)) error {
	return n.subscriber.Subscribe(msg, callback)
}

func (n *Node) Publish(msg proto.Message) error {
	return n.publisher.Publish(msg)
}
