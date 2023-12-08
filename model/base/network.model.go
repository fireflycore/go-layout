package base

type HostPortEntity struct {
	Host string `json:"host" bson:"host" yaml:"host" mapstructure:"host"`
	Port int    `json:"port" bson:"port" yaml:"port" mapstructure:"port"`
}

type ProtocolDomainNameEntity struct {
	Protocol   uint   `json:"protocol" bson:"protocol" yaml:"protocol" mapstructure:"protocol"`
	DomainName string `json:"domain_name" bson:"domain_name" yaml:"domain_name" mapstructure:"domain_name"`
}

type TlsEntity struct {
	CaCert         string `json:"ca_cert" bson:"ca_cert" yaml:"ca_cert" mapstructure:"ca_cert"`
	ServerCert     string `json:"server_cert" bson:"server_cert" yaml:"server_cert" mapstructure:"server_cert"`
	ServerCertKey  string `json:"server_cert_key" bson:"server_cert_key" yaml:"server_cert_key" mapstructure:"server_cert_key"`
	ClientCert     string `json:"client_cert" bson:"client_cert" yaml:"client_cert" mapstructure:"client_cert"`
	ClientCertKey  string `json:"client_cert_key" bson:"client_cert_key" yaml:"client_cert_key" mapstructure:"client_cert_key"`
	ClusterCert    string `json:"cluster_cert" bson:"cluster_cert" yaml:"cluster_cert" mapstructure:"cluster_cert"`
	ClusterCertKey string `json:"cluster_cert_key" bson:"cluster_cert_key" yaml:"cluster_cert_key" mapstructure:"cluster_cert_key"`
}

type NetworkAddress struct {
	Inside  string `json:"inside" yaml:"inside" mapstructure:"inside"`
	Outside string `json:"outside" yaml:"outside" mapstructure:"outside"`
}
