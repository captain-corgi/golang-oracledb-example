all: tidy vendor build run

build:
	go build -o app cmd/app/main.go

run:
	./app

clean:
	rm ./app
	rm -f vendor

tidy:
	go mod tidy

vendor:
	go mod vendor

.PHONY: coverage
# coverage:
# 	go test \
# 		-race -covermode=atomic -timeout 30s \
# 		-coverprofile=coverage/coverage.out \
# 		github.com/captain-corgi/golang-os-example/pkg/iptrans