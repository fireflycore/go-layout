package server

import (
	"go-layout/internal/conf"
	"net"
	"strconv"
)

func NewServer(bootstrapConf *conf.BootstrapConf) net.Listener {
	addr := net.JoinHostPort("0.0.0.0", strconv.FormatUint(uint64(bootstrapConf.Port), 10))
	listen, err := net.Listen("tcp", addr)

	if err != nil {
		panic(err)
	}

	return listen
}
