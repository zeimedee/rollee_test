download:
		go mod download 
		
run: 
	go run cmd/server/main.go
	
test:
	go test -v -cover ./...

run-docker:
	docker-compose up --build

stop-docker:
	docker-compose down