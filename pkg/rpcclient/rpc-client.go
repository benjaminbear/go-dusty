package rpcclient

import (
	"context"
	"time"

	"github.com/benjaminbear/go-dusty/pkg/protocol"

	"google.golang.org/grpc"
)

type rpcConn struct {
	connection *grpc.ClientConn
	client     protocol.DustyClient
	context    context.Context
	cancel     func()
}

type RpcConnection interface {
	Setup(username string, specsRepo string, vmMemory int32, sshKeyPath string, update bool) error
	Close() error
}

func CreateConnection() (RpcConnection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// address can be "unix:////var/run/dusty.sock" in unix
	conn, err := grpc.DialContext(ctx, protocol.Address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := protocol.NewDustyClient(conn)

	return &rpcConn{
		connection: conn,
		client:     client,
		context:    ctx,
		cancel:     cancel,
	}, nil
}

func (r *rpcConn) Close() error {
	return r.connection.Close()
}

func (r *rpcConn) Setup(username string, specsRepo string, vmMemory int32, sshKeyPath string, update bool) error {
	defer r.cancel()

	request := &protocol.SetupRequest{
		Username:   username,
		SpecsRepo:  specsRepo,
		VmMemory:   vmMemory,
		SshKeyPath: sshKeyPath,
		Update:     update,
	}

	reply, err := r.client.Setup(r.context, request)
	if err != nil {
		return err
	}

	return reply.HandleRegularReply()
}
