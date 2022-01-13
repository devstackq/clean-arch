package transport

type httpFactory struct{}

func (factory httpFactory) GetProtocol() Delivery {
	return &httpServer{}
}
