package config

type EndpointConfig interface {
	GetHost() string
	GetPort() int
}

type ServerConfig struct {
	Grpc         GrpcConfig    `fig:"grpc"`
	Gateway      GatewayConfig `fig:"gateway"`
	DnsHostnames []string      `fig:"dns_hostnames"`
}

type GrpcConfig struct {
	Host string `fig:"host" default:"localhost"`
	Port int    `fig:"port" default:"50051"`
}

func (c GrpcConfig) GetHost() string {
	return c.Host
}

func (c GrpcConfig) GetPort() int {
	return c.Port
}

type GatewayConfig struct {
	Host string `fig:"host" default:"localhost"`
	Port int    `fig:"port" default:"9090"`
}

func (c GatewayConfig) GetHost() string {
	return c.Host
}

func (c GatewayConfig) GetPort() int {
	return c.Port
}

type TlsConfig struct {
	UseTls   bool
	CaPath   string
	CertPath string
	KeyPath  string
}
