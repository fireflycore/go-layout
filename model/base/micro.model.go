package base

type MicroserviceEntity struct {
	Name      string `json:"name"`
	Endpoints string `json:"endpoints"`
}

type MicroserviceDiscoverEntity struct {
	Gateway         string `json:"gateway"`
	ServiceInstance string `json:"service_instance"`
	Service         string `json:"service"`
}
