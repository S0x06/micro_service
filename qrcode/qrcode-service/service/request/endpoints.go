package request

/*
import (
	"context"
	"fmt"

	"github.com/go-kit/kit/endpoint"
	// pb "github.com/_example"
)

type Endpoints struct {
	QrCodeEndpoint endpoint.Endpoint
}

func (e Endpoints) QrCode(ctx context.Context, in *pb.QrCodeRequest) (*pb.EchoResponse, error) {
	response, err := e.QrCodeEndpoint(ctx, in)
	if err != nil {
		return nil, err
	}
	return response.(*pb.QrCodeResponse), nil
}

func MakeQrCodeEndpoint(s pb.QrCodeServer) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*pb.QrCodeRequest)
		v, err := s.QrCode(ctx, req)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
*/
