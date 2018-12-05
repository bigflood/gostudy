package pb

//go:generate protoc -I . todo.proto --go_out=plugins=grpc:.
//go:generate flatc --proto --grpc   todo.proto
//go:generate flatc --grpc --go  todo.fbs
