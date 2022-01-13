package transport

type grpcFactory struct{}

func (factory grpcFactory) GetTransport() Transport {
	return &grpcServer{}
}
