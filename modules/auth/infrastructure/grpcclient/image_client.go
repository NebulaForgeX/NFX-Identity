package grpcclient

import (
	"nebulaid/pkgs/grpcx"
	"nebulaid/pkgs/security/token/servertoken"
	imagepb "nebulaid/protos/gen/image/image"

	"google.golang.org/grpc"
)

type ImageGRPCClient struct {
	conn      *grpc.ClientConn
	ImageStub imagepb.ImageServiceClient
}

func NewImageGRPCClient(addr string, provider servertoken.TokenProvider) (*ImageGRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpcx.DefaultClientOptions(provider)...)
	if err != nil {
		return nil, err
	}

	return &ImageGRPCClient{
		conn:      conn,
		ImageStub: imagepb.NewImageServiceClient(conn),
	}, nil
}

func (c *ImageGRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
