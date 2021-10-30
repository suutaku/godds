package tools

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func Hash(input []byte) string {
	hasher := sha1.New()
	hasher.Write(input)
	return hex.EncodeToString(hasher.Sum(nil))
}

func MessageId(input proto.Message) string {
	s := fmt.Sprintf("%s", input.ProtoReflect().Descriptor())
	return Hash([]byte(s))
}

func AnyFromString(input string) *anypb.Any {
	datas := strings.Split(input, ":")
	len := len(datas)
	if len != 2 {
		return nil
	}
	return &anypb.Any{
		TypeUrl: datas[0][1:len],
		Value:   []byte(datas[1]),
	}
}
