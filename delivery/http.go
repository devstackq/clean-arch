package delivery

import (
	"net/http"
	"time"

	handler "github.com/devstackq/go-clean/auth/delivery/http"
)

//init config http server
type httpServer struct {
	// authUseCase auth.UseCase
}

func NewHttpServer() httpServer {
	return httpServer{}
}

func (h httpServer) InitHttp(port string) http.Server {
	//init router
	hr := handler.NewHandler(app.authUseCase)
	http.HandleFunc("/signup", hr.SignUp) //register handler

	server := http.Server{
		Addr: ":" + port,
		// Handler: ,
		ReadTimeout:    10 * time.Second, // each 10 sec read
		WriteTimeout:   10 * time.Second, //each 10sec write
		MaxHeaderBytes: 1 << 20,          // max 20 mg
	}
	return server
}

func (httpServer) InitGrpc() {}
