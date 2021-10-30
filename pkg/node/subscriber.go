package node

import (
	"context"
	"sync"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type Subscriber struct {
	connector *Connector
	topics    map[string]func(msg *anypb.Any)
	running   bool
	mutex     sync.Mutex
}

func NewSubscriber(ctx context.Context, conn *Connector) *Subscriber {

	ret := &Subscriber{
		connector: conn,
		topics:    make(map[string]func(msg *anypb.Any)),
		running:   true,
	}
	return ret
}

func (s *Subscriber) Start() {
	for s.running {
		b, err := s.connector.Read()
		if err != nil {
			continue
		}
		var anyt = anypb.Any{}
		err = proto.Unmarshal(b, &anyt)
		if err != nil {
			continue
		}
		s.mutex.Lock()
		for k, v := range s.topics {
			if k == string(anyt.MessageName()) {
				v(&anyt)
			}
		}
		s.mutex.Unlock()
	}
}

func (s *Subscriber) Stop() {
	s.running = false
}

func (s *Subscriber) Subscribe(msg proto.Message, callback func(msg *anypb.Any)) error {
	s.mutex.Lock()
	s.topics[string(msg.ProtoReflect().Descriptor().FullName())] = callback
	s.mutex.Unlock()
	return nil
}

func (s *Subscriber) UnSubscribe(msg proto.Message) error {
	delete(s.topics, string(msg.ProtoReflect().Descriptor().FullName()))
	return nil
}
