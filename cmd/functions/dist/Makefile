build-DistFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap dist.go
	mv bootstrap $(ARTIFACTS_DIR)