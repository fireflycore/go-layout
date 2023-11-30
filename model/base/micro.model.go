package base

type MicroserviceEntity struct {
	Name      string `json:"name"`
	Endpoints string `json:"endpoints"`
}

type MicroserviceDiscoverEntity struct {
	Gateway   string `json:"gateway"`
	Namespace string `json:"namespace"`
	Service   string `json:"service"`
}
