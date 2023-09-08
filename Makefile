# load .env file

test:
	@echo "[Running test]"
	@go test -v internal/storage/*.go

