protoc --go_out=. --go-grpc_out=. hello.proto
cp hello.proto protos/
go build -o server server.go
go build -o client client.go
