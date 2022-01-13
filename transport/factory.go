package transport

type Factory interface {
	GetTransport() Transport
}

func GetFactory(typeProtocol string) Factory {
	if typeProtocol == "http" {
		return httpFactory{}
	} else if typeProtocol == "grpc" {
		return grpcFactory{}
	}
	// }else if typeProtocol == "mongo"
	return nil
}
