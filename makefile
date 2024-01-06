run:
	go run cmd/main/main.go

test:
	go test -v ./... 

build-docker:
	docker build -t tapi-controller .

docker-deploy:
	docker-compose up -d

docker-clear:
	docker-compose down -v