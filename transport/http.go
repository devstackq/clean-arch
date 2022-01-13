package transport

import (
	"net/http"
	"time"
)

type httpServer struct{}

func (h httpServer) InitTransport(port string) interface{} {

	return http.Server{
		Addr: ":" + port,
		// Handler:        authHttp.InitRoutes(usecase),
		ReadTimeout:    10 * time.Second, // each 10 sec read
		WriteTimeout:   10 * time.Second, //each 10sec write
		MaxHeaderBytes: 1 << 20,          // max 20 mg
	}
}
