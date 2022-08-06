build-PlushBotFunction:
	GOOS=linux go build -o bootstrap ./cmd
	cp ./bootstrap $(ARTIFACTS_DIR)/bootstrap
