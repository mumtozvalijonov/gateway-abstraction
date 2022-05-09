proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/server/proto/sum.proto

build:
	go build -o server cmd/main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run
