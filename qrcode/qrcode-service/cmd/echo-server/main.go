package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

type QrCodeService interface {
	GetQrCode(string) (string, error)
}

type qrCodeService struct{}

func (qrCodeService) GetQrCode(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return s, nil
}

var ErrEmpty = errors.New("empty string")

type getQrCodeRequest struct {
	S string `json:"s"`
}

type getQrCodeResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't define JSON marshaling
}

// Endpoints are a primary abstraction in go-kit. An endpoint represents a single RPC (method in our service interface)
func makeGetQrCodeEndpoint(svc QrCodeService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(getQrCodeRequest)
		v, err := svc.GetQrCode(req.S)
		if err != nil {
			return getQrCodeResponse{v, err.Error()}, nil
		}
		return getQrCodeResponse{v, ""}, nil
	}
}

func main() {

	// service.Run()
	svc := qrCodeService{}

	getQrCodeHandler := httptransport.NewServer(
		makeGetQrCodeEndpoint(svc),
		decodeGetQrCodeRequest,
		encodeResponse,
	)

	http.Handle("/getQrCode", getQrCodeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func decodeGetQrCodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getQrCodeRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
