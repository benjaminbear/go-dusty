package constants

const (
	ExampleRepo           = "github.com/gamechanger/dusty-example-specs"
	SocketAck             = `\1\1`
	SocketTerminator      = `\0\0`
	SocketErrorTerminator = `\0\1`
	SocketLoggerName      = "socket_logger"

	SocketPath       = "/var/run/dusty/dusty.sock"
	ContainerLogPath = "/var/log"
)
