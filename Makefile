BINARY_NAME=lker

gomod-init:
	go mod init github.com/ecshreve/${BINARY_NAME}

gomod-tidy:
	go mod tidy

go-build:
	go build -o bin/${BINARY_NAME} github.com/ecshreve/${BINARY_NAME}/cmd/${BINARY_NAME}

go-install:
	go install -i github.com/ecshreve/${BINARY_NAME}/cmd/${BINARY_NAME}

go-run: go-build
	bin/${BINARY_NAME}
	
go-test:
	go test github.com/ecshreve/${BINARY_NAME}/...

go-testv:
	go test -v github.com/ecshreve/${BINARY_NAME}/...

go-testc:
	go test -race -coverprofile=coverage.txt -covermode=atomic github.com/ecshreve/${BINARY_NAME}/...

go-all: test build run


docker-build:
	docker build --network="host" --tag registry.digitalocean.com/shreggie/lker:custom .

docker-run:
	docker run -d --network="host" --name="lker" registry.digitalocean.com/shreggie/lker:custom 

docker-run-prod:
	docker run -d --publish="0.0.0.0:80:8880 --name="lker" registry.digitalocean.com/shreggie/lker:custom 
