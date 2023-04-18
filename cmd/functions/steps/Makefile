build-StepsFunction:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap steps.go
	mv bootstrap $(ARTIFACTS_DIR)