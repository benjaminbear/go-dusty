package rpcserver

import (
	"context"
	"log"
	"net"

	"github.com/schlund/go-dusty/pkg/config"

	"github.com/schlund/go-dusty/pkg/protocol"

	"google.golang.org/grpc"
)

type rpcServer struct {
	protocol.UnimplementedDustyServer
}

func (s *rpcServer) Setup(ctx context.Context, in *protocol.SetupRequest) (*protocol.RegularReply, error) {
	reply := &protocol.RegularReply{}

	conf, err := config.LoadConfiguration()
	if err != nil {
		return reply.AddError(err), nil
	}

	conf.SetUsername(in.Username)
	conf.SetSpecsRepo(in.SpecsRepo)
	conf.SetVmMemorySize(int(in.VmMemory))
	conf.SetSetupHasRun()

	err = conf.SaveConfiguration()
	if err != nil {
		return reply.AddError(err), nil
	}

	// TODO: Implement warning system? (dusty original)

	if in.Update {
		reply.AddMessage("Pulling latest updates for all active managed repos:")
		// TODO: update spcecs and download repo
	}

	return reply.AddMessage("Initial setup completed. You should now be able to use Dusty!"), nil
}

func RunServer() {
	// Unix Sockets net.Listen("unix", "/tmp/socksock.sock")
	lis, err := net.Listen("tcp", protocol.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	protocol.RegisterDustyServer(s, &rpcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
