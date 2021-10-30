package test

import (
	"fmt"
	"testing"

	"github.com/suutaku/godds/pkg/tools"
	"github.com/suutaku/godds/proto"
)

func TestId(t *testing.T) {
	var msg = &proto.Version{}
	id := tools.MessageId(msg.ProtoReflect().Interface())
	fmt.Println(id)
}
