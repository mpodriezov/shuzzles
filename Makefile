run:
	@go run cmd/website/main.go
build:
	@go build -o bin/website cmd/website/main.go

up: 
	@miflo up

down: 
	@miflo revert
