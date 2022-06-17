if any protoc binary not found please use below command and then add that bin to $PATH.
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go

go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc


.proto have the whole path ention so run it with below command.
The .proto file is kept in root folder for same reason.
protoc --go-grpc_out=. ./usermgmt.proto --go_out=. ./usermgmt.proto

