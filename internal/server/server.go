package server

import (
	"fmt"
	"go-layout/internal/conf"
	"net"
)

func NewServer(bc *conf.BootstrapConf) net.Listener {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", bc.Port))

	if err != nil {
		panic(err)
	}

	return listen
}
