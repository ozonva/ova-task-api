rmdir /S /Q  pkg\api
mkdir vendor.protogen\api\ova-task-api\
copy api\ova-task-api\ova-task-api.proto vendor.protogen\api\ova-task-api\ova-task-api.proto
protoc -I vendor.protogen --go_out=pkg --go_opt=paths=source_relative  --go-grpc_out=pkg --go-grpc_opt=paths=source_relative  --grpc-gateway_out=pkg --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative api/ova-task-api/ova-task-api.proto