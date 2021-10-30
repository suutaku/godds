package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/suutaku/godds/pkg/node"
	"github.com/suutaku/godds/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RPCServer struct {
	proto.UnimplementedCoreServer
	node *node.Node
}

func NewRPCServer(node *node.Node) *RPCServer {
	return &RPCServer{
		node: node,
	}
}

func (s *RPCServer) Publish(ctx context.Context, input *anypb.Any) (*emptypb.Empty, error) {
	if input == nil {
		return &empty.Empty{}, errors.New("input is nil")
	}
	err := s.node.Publish(input)
	return &empty.Empty{}, err
}

func (s *RPCServer) Echo(ctx context.Context, input *anypb.Any) (*emptypb.Empty, error) {
	if input == nil {
		return &empty.Empty{}, errors.New("input is nil")
	}
	err := s.node.Subscribe(input, func(a *anypb.Any) {
		fmt.Printf("%+v", a)
	})
	return &empty.Empty{}, err
}
