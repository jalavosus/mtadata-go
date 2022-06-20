package env

const (
	hostKey string = "_HOST"
	portKey string = "_PORT"
)

const (
	PrefixDb   string = "DB"
	DbHost            = PrefixDb + hostKey
	DbPort            = PrefixDb + portKey
	DbUsername        = PrefixDb + "_USERNAME"
	DbPassword        = PrefixDb + "_PASSWORD"
	DbDatabase        = PrefixDb + "_DATABASE"
	DbSslMode         = PrefixDb + "_SSL_MODE"
)

const (
	PrefixGrpc string = "GRPC"
	GrpcHost          = PrefixGrpc + hostKey
	GrpcPort          = PrefixGrpc + portKey
)

const (
	PrefixGateway string = "GATEWAY"
	GatewayHost          = PrefixGateway + hostKey
	GatewayPort          = PrefixGateway + portKey
)