package transport

type Factory interface {
	GetTransport() Transport
}

// type FactoryDb interface {
// 	GetDatabase() Database
// }

func GetFactory(typeFactory string) Factory {
	if typeFactory == "http" {
		return httpFactory{}
	} else if typeFactory == "grpc" {
		return grpcFactory{}
	}
	// }else if typeFactory == "mongo"
	return nil
}
