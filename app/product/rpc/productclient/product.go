// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package productclient

import (
	"context"

	"mfzero-mall/app/product/rpc/product"

	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRequest  = product.CreateRequest
	CreateResponse = product.CreateResponse
	DetailRequest  = product.DetailRequest
	DetailResponse = product.DetailResponse
	RemoveRequest  = product.RemoveRequest
	RemoveResponse = product.RemoveResponse
	UpdateRequest  = product.UpdateRequest
	UpdateResponse = product.UpdateResponse

	Product interface {
		Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
		Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
		Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error)
		Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Create(ctx, in, opts...)
}

func (m *defaultProduct) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Update(ctx, in, opts...)
}

func (m *defaultProduct) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Remove(ctx, in, opts...)
}

func (m *defaultProduct) Detail(ctx context.Context, in *DetailRequest, opts ...grpc.CallOption) (*DetailResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.Detail(ctx, in, opts...)
}
