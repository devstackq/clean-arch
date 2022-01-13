package transport

type grpcServer struct{}

func (grpcServer) InitTransport(port string) interface{} {
	// app.grpcServer = *grpc.NewServer()
	// list, err := net.Listen("tcp", ":8000")
	// if err != nil {
	// 	fmt.Println("SOMETHING HAPPEN")
	// }
	// app.grpcServer.Serve(list)
	return nil
}
