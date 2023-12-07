package utils

import (
	"encoding/json"
	"fmt"
	"microservice-go/utils/mongo"
	"microservice-go/utils/time"
)

type _Entrance struct {
	Time  time.Entrance
	Mongo mongo.Entrance
}

var Use = new(_Entrance)

func StructConvert[S any, T any](source *S, target *T) {
	marshal, mErr := json.Marshal(source)
	if mErr != nil {
		fmt.Println(mErr.Error())
		return
	}
	if err := json.Unmarshal(marshal, &target); err != nil {
		fmt.Println(err.Error())
		return
	}
}
