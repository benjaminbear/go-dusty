package constants

import (
	"os"
	"path/filepath"
)

const (
	ExampleRepo = "github.com/gamechanger/dusty-example-specs"

	SocketPath       = "/var/run/dusty/dusty.sock"
	ContainerLogPath = "/var/log"
)

var (
	DefaultSSHKeyPath     = filepath.Join(os.Getenv("HOME"), "/.ssh/id_rsa")
	DefaultKnownHostsPath = filepath.Join(os.Getenv("HOME"), "/.ssh/known_hosts")
)
