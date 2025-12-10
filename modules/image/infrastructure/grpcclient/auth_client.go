package grpcclient

import (
	"nebulaid/pkgs/grpcx"
	"nebulaid/pkgs/security/token/servertoken"
	profilepb "nebulaid/protos/gen/auth/profile"
	userpb "nebulaid/protos/gen/auth/user"

	"google.golang.org/grpc"
)

type AuthGRPCClient struct {
	conn        *grpc.ClientConn
	UserStub    userpb.UserServiceClient
	ProfileStub profilepb.ProfileServiceClient
}

func NewAuthGRPCClient(addr string, provider servertoken.TokenProvider) (*AuthGRPCClient, error) {
	conn, err := grpc.NewClient(addr, grpcx.DefaultClientOptions(provider)...)
	if err != nil {
		return nil, err
	}

	return &AuthGRPCClient{
		conn:        conn,
		UserStub:    userpb.NewUserServiceClient(conn),
		ProfileStub: profilepb.NewProfileServiceClient(conn),
	}, nil
}

func (c *AuthGRPCClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
