package transport

type Delivery interface {
	InitTransport(port string) interface{}
}
