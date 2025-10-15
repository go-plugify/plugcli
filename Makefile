build:
	rm -rf build
	mkdir -p build
	go build -o build/plugcli ./

.PHONY: build