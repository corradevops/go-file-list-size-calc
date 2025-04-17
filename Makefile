BINARY_NAME=filesizecalc

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-arm64
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux-amd64
	GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}-linux-arm64
	GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows.exe

run: build
	./${BINARY_NAME}-linux

clean:
	go clean
	@if [ -f "${BINARY_NAME}-darwin-amd64" ]; then \
		rm ${BINARY_NAME}-darwin-amd64; \
	fi
	@if [ -f "${BINARY_NAME}-darwin-arm64" ]; then \
		rm ${BINARY_NAME}-darwin-arm64; \
	fi
	@if [ -f "${BINARY_NAME}-linux-amd64" ]; then \
		rm ${BINARY_NAME}-linux-amd64; \
	fi
	@if [ -f "${BINARY_NAME}-linux-arm64" ]; then \
		rm ${BINARY_NAME}-linux-arm64; \
	fi
	@if [ -f "${BINARY_NAME}-windows.exe" ]; then \
		rm "${BINARY_NAME}-windows.exe"; \
	fi
