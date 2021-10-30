package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	cli "github.com/jawher/mow.cli"
	mycli "github.com/suutaku/godds/internal/cli"
	"github.com/suutaku/godds/pkg/node"
	"github.com/suutaku/godds/pkg/tools"
	"github.com/suutaku/godds/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
)

var n *node.Node
var ctx context.Context
var port = "5008"
var addr = "127.0.0.1"

func cmdDaemon(cmd *cli.Cmd) {
	cmd.Action = func() {
		n.Start()
		s := grpc.NewServer()
		proto.RegisterCoreServer(s, mycli.NewRPCServer(n)) // ここでembedding済のserverを渡す

		lis, err := net.Listen("tcp", addr+":"+port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}

	}
}

func cmdSubscrib(cmd *cli.Cmd) {
	cmd.Spec = "TOPIC"
	topic := cmd.StringArg("TOPIC", "", "Specifiy a topic")
	if topic == nil {
		return
	}

	cmd.Action = func() {
		conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		client := proto.NewCoreClient(conn)
		ctxs, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		var input = tools.AnyFromString(*topic)

		if input == nil {
			fmt.Println("invalid value")
			return
		}
		_, err = client.Echo(ctxs, input)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func cmdPublish(cmd *cli.Cmd) {
	exmple, _ := anypb.New(&proto.Version{})
	cmd.Spec = "[ -t=<type> ] VALUE"
	dType := cmd.StringOpt("t type", "", "data type [json,xml,raw...]")
	dValue := cmd.StringArg("VALUE", "", "data value, like"+fmt.Sprintf("%v", exmple))
	if dType == nil || dValue == nil {
		return
	}
	cmd.Action = func() {
		fmt.Printf("data type is %s, value is %s\n", *dType, *dValue)
		conn, err := grpc.Dial(addr+":"+port, grpc.WithInsecure())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer conn.Close()
		client := proto.NewCoreClient(conn)
		ctxs, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		var input = tools.AnyFromString(*dValue)

		if input == nil {
			fmt.Println("invalid value")
			return
		}
		_, err = client.Publish(ctxs, input)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ctx = context.Background()
	n = node.NewNode(ctx, "", "")
	app := cli.App("godds", "A simple dds service")

	app.Command("daemon", "Start dds daemon", cmdDaemon)
	app.Command("pub", "publish a topic", cmdPublish)
	app.Command("echo", "echo a topic", cmdSubscrib)

	app.Run(os.Args)

}
