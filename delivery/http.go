package delivery

import (
	"net/http"
	"time"
)

type server struct{}

func NewHttpServer() server {
	return server{}
}

func (h server) InitHTTP(port string) http.Server {

	return http.Server{
		Addr: ":" + port,
		// Handler: ,
		ReadTimeout:    10 * time.Second, // each 10 sec read
		WriteTimeout:   10 * time.Second, //each 10sec write
		MaxHeaderBytes: 1 << 20,          // max 20 mg
	}
}

func (server) InitGrpc() {}
