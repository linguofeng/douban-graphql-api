run:
	go run main.go

build:
	go build -o douban-graphql-api main.go

build-lambda:
	mkdir -p functions
	GO111MODULE=on go build -ldflags '-X main.isLambda=true' -o functions/graphql main.go