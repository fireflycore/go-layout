package dep

import (
	"go-layout/internal/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRemoteServiceGrpcClient(bc *conf.BootstrapConf) (*grpc.ClientConn, error) {
	var addr string
	if bc.Gateway.Network.SN == bc.Micro.Network.SN {
		addr = bc.Gateway.Network.Internal
	} else {
		addr = bc.Gateway.Network.External
	}
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
