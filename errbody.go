package errbody

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
)

type ErrorBody struct {
	Error string `protobuf:"bytes,1,name=error" json:"error"`
	// This is to make the error more compatible with users that expect errors to be Status objects:
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	// It should be the exact same message as the Error field.
	Message string     `protobuf:"bytes,1,name=message" json:"message"`
	Code    int32      `protobuf:"varint,2,name=code" json:"code"`
	Details []*any.Any `protobuf:"bytes,3,rep,name=details" json:"details,omitempty"`
}

// Make this also conform to proto.Message for builtin JSONPb Marshaler
func (e *ErrorBody) Reset()         { *e = ErrorBody{} }
func (e *ErrorBody) String() string { return proto.CompactTextString(e) }
func (*ErrorBody) ProtoMessage()    {}
