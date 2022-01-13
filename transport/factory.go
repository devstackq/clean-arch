package transport

type Factory interface {
	GetProtocol() Delivery
}

func GetFactory(typeProtocol string) Factory {
	if typeProtocol == "http" {
		return httpFactory{}
	} else if typeProtocol == "grpc" {
		return grpcFactory{}
	}
	return nil
}
