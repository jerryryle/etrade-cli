all: test build

test:
	go test ./...

build:
	go build -C ./etrade

install:
	@printf "Installing to: "
	@go list -C ./etrade -f '{{.Target}}'
	@go install -ldflags "-s -w" ./etrade
	@echo "Done!"

clean:
	go clean
	rm -f ./etrade/etrade
