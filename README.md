# golang-grpc-demo-initial

golang-grpc-demo-initial

# to generate proto files

run protoc --go_out=. --go-grpc_out=. proto/greet.proto

or

protoc --go_out=. --go_opt=module=github.com/leulad/golang-grpc-demo-initial --go-grpc_out=. --go-grpc_opt=module=github.com/leulad/golang-grpc-demo-initial/greet.proto
