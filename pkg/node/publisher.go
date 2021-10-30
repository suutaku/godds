package node

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Publisher struct {
	connector *Connector
}

func NewPublisher(conn *Connector) *Publisher {
	return &Publisher{
		connector: conn,
	}
}

func (p *Publisher) Publish(data proto.Message) error {
	anyt, err := anypb.New(data)
	if err != nil {
		return err
	}
	b, err := proto.Marshal(anyt)
	if err != nil {
		return err
	}
	return p.connector.Write(b)
}

func (p *Publisher) Start() {
	p.connector.Start()
}
