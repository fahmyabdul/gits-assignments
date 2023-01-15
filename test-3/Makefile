current_dir = $(shell pwd)
docker_repo = sir-avdul.ddns.net:8555/etcetera/test-3-grpc:latest

proto-gen-old:
	protoc -Iinternal/proto --go_out=. --go_opt=module=github.com/fahmyabdul/gits-assignments/test-3 --go-grpc_out=paths=source_relative:internal/proto internal/proto/*.proto
.PHONY: proto-gen-old

proto-gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/proto/*.proto
.PHONY: proto-gen

swagg-gen:
	swag init -g cmd/test-3-restapi/main.go -o api/open_api
.PHONY: swagg-gen

build-grpc:
	go mod tidy
	mkdir -p bin/
	bash scripts/build.sh --app test-3-grpc --out ./bin/
.PHONY: build-grpc

run-grpc:
	go mod tidy
	bash scripts/run.sh --app test-3-grpc
.PHONY: run

run-grpc-dev:
	go mod tidy
	bash scripts/run.sh --app test-3-grpc --config .config.dev.yml
.PHONY: run

gen-mocks:
	mockery --dir=./internal/usecase --all -r --case snake --output internal/usecase/mocks
.PHONY: gen-mocks

export TEST_HOST=localhost:8500

test:
	go clean -testcache && go test -cover ./...
.PHONY: test

test-verbose:
	go clean -testcache && go test -cover ./... -v
.PHONY: test

test-integration:
	go clean -testcache && TEST_INTEGRATION=true go test -v ./integration-test
.PHONY: test-integration

clean:
	rm -rf log/*
	rm -rf tmp/*
.PHONY: clean

docker-build:
	docker build -t ${docker_repo} -f Dockerfile .
.PHONY: docker-build

docker-run:
	docker-compose -f deployments/compose/docker-compose.yml up
.PHONY: docker-run

docker-run-detached:
	docker-compose -f deployments/compose/docker-compose.yml up -d
.PHONY: docker-run-detached

docker-logs:
	docker logs -f -n 100 test-3-grpc
.PHONY: docker-logs

docker-stop:
	docker-compose -f deployments/compose/docker-compose.yml down
.PHONY: docker-stop

docker-build-push:
	docker build -t ${docker_repo} -f Dockerfile .
	docker push ${docker_repo}
.PHONY: docker-build-push