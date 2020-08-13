gen-cal:
	 protoc demopb/demo.proto --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.

run-server:
	go run main.go

run-proxy:
	go run proxy/proxy.go