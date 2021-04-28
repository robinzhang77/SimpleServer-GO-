.PHONY: init
init: ## Init project config
	@git config core.hooksPath githooks

.PHONY: proto
proto: ## Complie proto file
	./protoc.exe --gogofaster_out=./pb ./proto/*.proto --proto_path=./proto/

.PHONY: submodule
submodule: ## Update submodule
	git submodule update --init

.PHONY: build
build: ## Build backend and client
	go build -o main.exe ./main.go

.PHONY: run
run: ## Startup backend with dev config
	go run ./main.go

