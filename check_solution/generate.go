package send_solution

// чтоб работало c gogoproto
// go get github.com/gogo/protobuf/proto
// go get github.com/gogo/protobuf/jsonpb
// go get github.com/gogo/protobuf/protoc-gen-gogo
// go get github.com/gogo/protobuf/gogoproto
//go:generate protoc -I=. -I=$GOPATH/src/github.com/gogo/protobuf/protobuf -I=$GOPATH/src --micro_out=. --go_out=plugins=grpc:. check_solution.proto