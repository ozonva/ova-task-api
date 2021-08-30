.PHONY: build
build:
	go build -o $(CURDIR)/bin $(CURDIR)/cmd/ova-task-api/main.go

.PHONY: generate
generate:
	rm -r -f pkg/api
	mkdir -p vendor.protogen/api/ova-task-api/
	cp api/ova-task-api/ova-task-api.proto vendor.protogen/api/ova-task-api/ova-task-api.proto
	protoc -I vendor.protogen \
            --go_out=pkg --go_opt=paths=source_relative  \
            --go-grpc_out=pkg --go-grpc_opt=paths=source_relative  \
			--grpc-gateway_out=pkg \
			--grpc-gateway_opt=logtostderr=true \
			--grpc-gateway_opt=paths=source_relative \
            api/ova-task-api/ova-task-api.proto

.PHONY: generate-win
generate-win:
	generate-proto.bat


.PHONY: .install-go-deps
.install-go-deps:
	ls go.mod || go mod init
	go get github.com/golang/protobuf/proto
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/grpc-ecosystem/grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go get google.golang.org/grpc
	go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get google.golang.org/protobuf/cmd/protoc-gen-go
	go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/protobuf/cmd/protoc-gen-go

.PHONY: generate-swagger
generate-swagger:
	rm -r -f swagger
	mkdir -p swagger
	protoc -I vendor.protogen --swagger_out=allow_merge=true,merge_file_name=api:swagger api/ova-task-api/ova-task-api.proto

.PHONY: run-swagger
run-swagger: generate-swagger
	docker run --name swagger-container -p 80:8080 -e BASE_URL=/swagger -e SWAGGER_JSON=/swagger/api.swagger.json -v `pwd`/swagger:/swagger swaggerapi/swagger-ui
