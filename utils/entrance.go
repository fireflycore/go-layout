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

func Bool(val int) bool {
	var state bool
	switch val {
	case 0:
		state = false
	case 1:
		state = true
	}
	return state
}
