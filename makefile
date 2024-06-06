.PHONY: default run build test docs clean

APP_NAME=maybe

default=run

run: 
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:	
	@rm -f $(APP_NAME)