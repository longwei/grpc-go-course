# gRPC_playground
 
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/longwei/grpc-go-course --go-grpc_out=. --go-grpc_opt=module=github.com/longwei/grpc-go-course greet/proto/dummy.proto 