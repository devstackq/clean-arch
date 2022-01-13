package transport

type Transport interface {
	InitTransport(port string) interface{}
}
