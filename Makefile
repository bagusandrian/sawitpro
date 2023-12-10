

.PHONY: build

# build
build: 
	@echo " > Building [sawit-pro]..."
	@cd ./cmd && go build -o ../bin/sawit-pro && cd ../ 
	@echo "success build binary sawit pro bin/sawit-pro"

run:
	@echo " > Running [sawit-pro]..."
	@cd ./cmd && go run main.go && cd ../ 


init:
	go mod tidy
	go mod vendor

run-docker: 
	docker-compose up --build
test:
	go test -short -coverprofile coverage.out -v ./...