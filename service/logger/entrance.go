package logger

import (
	"context"
	"fmt"
	pb "microservice-go/gen/acme/logger/v1"
	"microservice-go/micro"
	"microservice-go/store"
	"strconv"
	"strings"
	"time"
)

type Entrance struct {
}

func (Entrance) Add(level uint32, message string) {
	endpoint := store.Use.Micro.Service["/microservice/lhdht/logger/Add"]

	if len(endpoint) == 0 {
		store.Use.Logger.Temp = append(store.Use.Logger.Temp, fmt.Sprintf("%dLPS%s", level, message))
		return
	}

	conn := micro.Use.Grpc.Dial(endpoint)
	if conn == nil {
		return
	}

	cli := pb.NewLoggerServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if len(store.Use.Logger.Temp) != 0 {
		for _, logs := range store.Use.Logger.Temp {
			str := strings.Split(logs, "LPS")
			l, _ := strconv.Atoi(str[0])

			if _, err := cli.Add(ctx, &pb.AddRequest{AppId: store.Use.Config.Local.System.AppId, Level: uint32(l), Message: str[1]}); err != nil {
				fmt.Println(err)
				return
			}
		}
		store.Use.Logger.Temp = nil
	}
	if _, err := cli.Add(ctx, &pb.AddRequest{AppId: store.Use.Config.Local.System.AppId, Level: level, Message: message}); err != nil {
		fmt.Println(err)
		return
	}
}
