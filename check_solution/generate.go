package check_solution

////go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --micro_out=. --go_out=plugins=grpc:. check_solution.proto

//go:generate protoc --go_out=plugins=grpc:. *.proto