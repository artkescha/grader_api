package grader_api

//// Not Working!
////go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --micro_out=. --go_out=plugins=grpc:. check_solution.proto

////Working!
///go:generate protoc --go_out=:. send_solution.proto
///go:generate protoc --go-grpc_out=. check_solution.proto
