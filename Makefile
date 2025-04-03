.PHONY: gen-client
	protoc --go_out=./mainservice/ --go-grpc_out=require_unimplemented_servers=false:./mainservice/  ./bookingservice/api/proto/booking.proto
	protoc --go_out=./mainservice/ --go-grpc_out=require_unimplemented_servers=false:./mainservice/  ./authorizationservice/proto/auth.proto


.PHONY: gen-server
	protoc --go_out=./bookingservice/ --go-grpc_out=./bookingservice/  ../bookingservice/api/proto/booking.proto
	protoc --go_out=./authorizationservice/ --go-grpc_out=./authorizationservice/  ../authorizationservice/proto/auth.proto
