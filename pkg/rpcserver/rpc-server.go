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
	log.Println("Setup called")

	conf, err := config.LoadConfiguration()
	if err != nil {
		return protocol.CreateRegularErrorReply(err), nil
	}

	conf.SetUsername(in.Username)
	conf.SetSpecsRepo(in.SpecsRepo)
	conf.SetVmMemorySize(int(in.VmMemory))
	conf.SetSetupHasRun()

	err = conf.SaveConfiguration()
	if err != nil {
		return protocol.CreateRegularErrorReply(err), nil
	}

	return &protocol.RegularReply{Message: "Successfully added config"}, nil
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
