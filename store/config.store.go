package store

import "microservice-go/config"

type _ConfigStoreEntrance struct {
	Local  *config.Entrance // local config
	Remote interface{}      // remote config
}
