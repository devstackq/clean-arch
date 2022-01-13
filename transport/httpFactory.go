package transport

type httpFactory struct{}

func (factory httpFactory) GetTransport() Transport {
	return &httpServer{}
}
