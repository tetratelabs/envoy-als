package sink

import (
	"fmt"
	"io"
	"log"

	v3 "github.com/envoyproxy/go-control-plane/envoy/service/accesslog/v3"
	"github.com/gogo/protobuf/jsonpb"
)

type server struct {
	marshaler jsonpb.Marshaler
}

var _ v3.AccessLogServiceServer = &server{}

func New() v3.AccessLogServiceServer {
	fmt.Println("Creating new ALS server")
	return &server{}
}

func (s *server) StreamAccessLogs(stream v3.AccessLogService_StreamAccessLogsServer) error {
	for {
		in, err := stream.Recv()
		log.Println("Received value")
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		str, _ := s.marshaler.MarshalToString(in)
		log.Println(str)
	}
}
