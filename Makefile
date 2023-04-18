.PHONY: build

aws:
	sam build

imaging-cli:
	go build -o ./bin/imaging-cli ./cmd/imaging-cli/main.go

build-DistFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/functions/dist/dist.go
	mv bootstrap $(ARTIFACTS_DIR)

build-StepsFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap ./cmd/functions/steps/steps.go
	mv bootstrap $(ARTIFACTS_DIR)