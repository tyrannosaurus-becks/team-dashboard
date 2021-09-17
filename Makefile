setup:
	go mod vendor

fmt:
	go fmt ./...

test:
	go test -v ./...

dev:
	cd cmd/send-metrics && go build && go install && cd -
