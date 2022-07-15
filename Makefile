build:
	go build -o out

run: build
	./out/metriq-rpc -config config/config.toml

