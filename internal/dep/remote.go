package dep

import (
	"go-layout/internal/conf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewRemoteServiceGrpcClient(bootstrapConf *conf.BootstrapConf) (*grpc.ClientConn, error) {
	var addr string
	if bootstrapConf.Gateway.Network.SN == bootstrapConf.Micro.Network.SN {
		addr = bootstrapConf.Gateway.Network.Internal
	} else {
		addr = bootstrapConf.Gateway.Network.External
	}
	return grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
