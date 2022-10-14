package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient will return 3 params is GRPCClient instance, CloseFunc, error.
// T must be particular GRPC Service Client interface. Like : ExampleServiceClient, HealthServiceClient...
// Example:
//
// serverAddr := "localhost:10443"
//
//	client, closeFunc, err := NewClient(serverAddr, func(conn grpc.ClientConnInterface) examplev1.ExampleServiceClient {
//		 return examplev1.NewExampleServiceClient(conn)
//	})
func NewClient[T any](serverAddr string, newClientFunc func(conn grpc.ClientConnInterface) T, opts ...grpc.DialOption) (T, func() error, error) {
	if len(opts) == 0 {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}
	var client T
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return client, nil, err
	}
	client = newClientFunc(conn)
	return client, conn.Close, err
}
