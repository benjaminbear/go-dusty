Build Instructions

`protoc -I pkg\protocol\ pkg\protocol\rpc.proto --go_out=plugins=grpc:pkg/protocol`