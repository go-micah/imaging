.PHONY: build

aws:
	sam build

imaging-cli:
	go build -o ./bin/imaging-cli ./cmd/imaging-cli/main.go