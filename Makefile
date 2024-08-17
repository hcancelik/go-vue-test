BINARY_NAME=test

.PHONY: clean
clean:
	rm -rf .${BINARY_NAME}
	rm -f ${BINARY_NAME}
	go clean

.PHONY: build
build: clean
	cd frontend && npm run build
	cd ..
	go mod tidy
	go build -tags prod -o ${BINARY_NAME} .
	chmod +x ${BINARY_NAME}
