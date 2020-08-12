gen-cal:
	protoc calculatorpb/calculator.proto --go_out=plugins=grpc:.

run-server:
	go run server/main.go

run-client:
	go run client/main.go