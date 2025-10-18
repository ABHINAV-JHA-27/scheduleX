.PHONY: clean  dev build

clean:
	@echo cleaning up project...
	@go mod tidy
	@echo dependencies cleaned successfully!!

dev:
	@echo running in dev...
	@CGO_ENABLED=1 go run cmd/scheduleX/main.go

build:
	@echo building...
	@CGO_ENABLED=1 go build cmd/scheduleX/main.go
	@echo application build successfully!!