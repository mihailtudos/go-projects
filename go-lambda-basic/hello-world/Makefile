build-HelloWorldFunction:
	GOOS=linux GOARCH=amd64 go build -o bootstrap
	chmod +x bootstrap
	cp ./bootstrap $(ARTIFACTS_DIR)/.