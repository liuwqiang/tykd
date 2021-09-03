VERSION=$(shell git describe --tags --abbrev=0)

rebuild:
	go clean -i
	go build

grpc:
	protoc -I . polit.proto --go_out=plugins=grpc:.

generate:
	./gormgen.sh 122.115.55.3:32453 root wftest@231 gateway *