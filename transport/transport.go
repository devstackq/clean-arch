package transport

type Transport interface {
	InitTransport(port string) interface{}
	// Selva()
}
//good practice 1 
// type Database interface {
// 	InitDb() (interface{}, error)
// }
