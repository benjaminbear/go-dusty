package rpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"path/filepath"

	"github.com/benjaminbear/go-dusty/constants"

	"github.com/benjaminbear/go-dusty/pkg/sshhandler"

	"github.com/benjaminbear/go-dusty/pkg/config"
	"gopkg.in/src-d/go-git.v4"

	"github.com/benjaminbear/go-dusty/pkg/protocol"

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
	conf.SetSSHKeyPath(in.SshKeyPath)
	conf.SetSetupHasRun()

	err = conf.SaveConfiguration()
	if err != nil {
		return reply.AddError(err), nil
	}

	// TODO: Implement warning system? (dusty original)

	if in.Update {
		reply.AddMessage("Pulling latest updates for all active managed repos:")

		repoPath, err := defaultReposPath()
		if err != nil {
			return reply.AddError(err), nil
		}

		parsedUrl, err := sshhandler.ParseUrl(conf.GetSpecsRepo())
		if err != nil {
			return reply.AddError(err), nil
		}

		opts := &git.CloneOptions{
			URL: parsedUrl.String(),
		}

		if parsedUrl.Scheme == "ssh" {
			// TODO: update specs and download repo
			opts.Auth, err = sshhandler.ParsePrivateKey(conf.GetSSHKeyPath())
			if err != nil {
				return reply.AddError(err), nil
			}

			err := sshhandler.AddHostIfMissing(parsedUrl.Host, constants.DefaultKnownHostsPath)
			if err != nil {
				return reply.AddError(err), nil
			}
		}

		path, err := sshhandler.BuildRepoPath(parsedUrl)
		if err != nil {
			return reply.AddError(err), nil
		}

		_, err = git.PlainClone(filepath.Join(repoPath, path), false, opts)
		if err != nil {
			return reply.AddError(err), nil
		}

		reply.AddMessage(fmt.Sprintf("Updating local repo %s", conf.GetSpecsRepo()))
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
