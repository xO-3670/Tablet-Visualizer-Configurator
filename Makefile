.PHONY: build

build:
	go build -ldflags="-s -w -H=windowsgui" -o Configurator.exe .
