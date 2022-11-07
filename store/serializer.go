package store

import (
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/proto"
)

var serializerMap = map[string]Serializer{
	"json":  &jsonSerializer{},
	"proto": &protoSerializer{},
}

func RegisterSerializer(name string, s Serializer) {
	serializerMap[name] = s
}

type Serializer interface {
	Marshal(interface{}) ([]byte, error)
	Unmarshal([]byte, interface{}) error
}

type jsonSerializer struct{}

func (j *jsonSerializer) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j *jsonSerializer) Unmarshal(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

type protoSerializer struct{}

func (p *protoSerializer) Marshal(v interface{}) ([]byte, error) {
	m, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("proto marshal value not proto.Message type but %T", v)
	}
	return proto.Marshal(m)
}

func (p *protoSerializer) Unmarshal(d []byte, v interface{}) error {
	m, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("proto unmarshal value not proto.Message type but %T", v)
	}
	return proto.Unmarshal(d, m)
}
