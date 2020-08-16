gen-cal:
	 protoc demopb/demo.proto --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true,paths=source_relative:.

run-server:
	go run main.go

run-proxy:
	go run proxy/proxy.go

proto:
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--go_out=plugins=grpc:. \
		demopb/demo.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--grpc-gateway_out=logtostderr=true:. \
		demopb/demo.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--swagger_out=logtostderr=true:. \
		demopb/demo.proto
	protoc -I/usr/local/include -I. \
		-I${GOPATH}/src \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
		--validate_out="lang=go:." \
		demopb/demo.proto