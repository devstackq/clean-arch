package delivery

type grpc struct{}

func NewGrpcServer() grpc {
	return grpc{}
}

func (grpc) InitHttp() {

}

func (grpc) InitGrpc() {
	// app.grpcServer = *grpc.NewServer()
	// list, err := net.Listen("tcp", ":8000")
	// if err != nil {
	// 	fmt.Println("SOMETHING HAPPEN")
	// }
	// app.grpcServer.Serve(list)
}
