package transport

type grpcFactory struct{}

func (factory grpcFactory) GetProtocol() Delivery {
	return &grpcServer{}
}
