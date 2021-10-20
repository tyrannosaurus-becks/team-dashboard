setup:
	go mod vendor

fmt:
	go fmt ./...

test:
	go test -v ./...

dev:
	cd cmd/platform-stats && go build && go install && cd -
	cd cmd/send-metrics && go build && go install && cd -
