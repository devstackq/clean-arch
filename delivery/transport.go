package delivery

import "net/http"

type Delivery interface {
	InitHttp(port string) http.Server
	InitGrpc()
}
