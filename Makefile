gen:
	protoc --proto_path=pb pb/*.proto --go_out=plugins=grpc:./src
clean:
	rm -r src/eci
run:
	go run ./src/server
