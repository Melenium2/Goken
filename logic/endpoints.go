package logic

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// >> Things represented below is auto-generated by KEN Generator <<

// Endpoints from the service interface.
type HelloEndpoints struct {
	Method1Endpoint  endpoint.Endpoint
	Methods3Endpoint endpoint.Endpoint
}

// Implementations of service for grpc client
func (e HelloEndpoints) Method1() error {
	panic("implement me pls...")
}
func (e HelloEndpoints) methods2(err error) error {
	panic("implement me pls...")
}
func (e HelloEndpoints) Methods3(r interface{}) interface{} {
	panic("implement me pls...")
}

// The functions that creates endpoints.
func MakeMethod1Endpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		panic("Implement me!!")
	}
}
func MakeMethods3Endpoint(s HelloService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (interface{}, error) {
		panic("Implement me!!")
	}
}

func NewHelloEndpoints(s HelloService) *HelloEndpoints {
	return &HelloEndpoints{
		Method1Endpoint:  MakeMethod1Endpoint(s),
		Methods3Endpoint: MakeMethods3Endpoint(s),
	}
}
