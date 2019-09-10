run:
	go run main.go

build:
	go build -o douban-graphql-api main.go

build-lambda:
	mkdir -p functions
	go get ./...
	go build -o functions/graphql main_lambda.go