package config

type EndpointConfig interface {
	GetHost() string
	GetPort() int
}

type ServerConfig struct {
	Grpc         HostPortConfig `fig:"grpc"`
	Gateway      HostPortConfig `fig:"gateway"`
	Graphql      HostPortConfig `fig:"graphql"`
	DnsHostnames []string       `fig:"dns_hostnames"`
}

type HostPortConfig struct {
	Host string `fig:"host" default:"localhost"`
	Port int    `fig:"port"`
}

func (c HostPortConfig) GetHost() string {
	return c.Host
}

func (c HostPortConfig) GetPort() int {
	return c.Port
}

type TlsConfig struct {
	UseTls   bool
	CaPath   string
	CertPath string
	KeyPath  string
}
