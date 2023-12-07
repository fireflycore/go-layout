package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"microservice-go/store"
	"strings"
	"time"
)

type Entrance struct {
}

func (Entrance) Dial(endpoint []string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var index int
	length := len(endpoint)
	if length == 0 {
		fmt.Println("no service endpoint are available")
		return nil
	} else if length == 1 {
		index = 0
	} else {
		// Todo 根据负载均衡算法去进行多节点服务负载
	}

	address := endpoint[index]
	switch store.Use.Config.Local.Micro.Deploy {
	case "LAN":
		srv := strings.Split(endpoint[index], ":")
		cur := strings.Split(store.Use.Config.Local.Micro.Address.Outside, ":")

		if srv[0] == cur[0] {
			address = strings.Join([]string{"127.0.0.1", srv[1]}, ":")
		}
	case "NAN":
		// not area network
		break
	}

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(fmt.Sprintf("the service endpoint is unavailable, error: %s", err.Error()))
		return nil
	}

	return conn
}
