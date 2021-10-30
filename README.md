# godds

A simple implementation of Data Distribution Service (DDS) in Go.
 
## example

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/suutaku/godds/pkg/node"
	"github.com/suutaku/godds/proto"
	gproto "google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	ctx := context.Background()
	node := node.NewNode(ctx, "", "")
	var msg = proto.Version{}
	node.Subscribe(&msg, func(data *anypb.Any) {
		var rmsg = proto.Message{}
		anypb.UnmarshalTo(data, &rmsg, gproto.UnmarshalOptions{})
		fmt.Println(data)
	})
	node.Start()

	for {
		sub := proto.Version{Version: "0.0.1"}
		node.Publish(&sub)
		time.Sleep(1 * time.Second)
	}

}
```