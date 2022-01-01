package delivery

import "net/http"

type Delivery interface {
	InitHTTP(port string) http.Server
	InitGrpc()
}
