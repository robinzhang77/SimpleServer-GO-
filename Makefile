.PHONY: init
init: ## Init project config
	@git config core.hooksPath githooks

.PHONY: proto
proto: ## Complie proto file
#	./protoc_client.exe --gogofaster_out=./pb ./proto/*.proto --proto_path=./proto/
#	./protoc_client.exe --go_out=./pb ./proto/*.proto --proto_path=./proto/
	./protoc_golang  --gogofaster_out=./pb ./proto/*.proto
	./protoc_client.exe ./proto/*.proto --csharp_out=S:/work/zgame/nvmeshdemo/Assets/Scripts/pb/

.PHONY: submodule
submodule: ## Update submodule
	git submodule update --init

.PHONY: build
build: ## Build backend and client
	go build -o main.exe ./main.go

.PHONY: run
run: ## Startup backend with dev config
	go run ./main.go

