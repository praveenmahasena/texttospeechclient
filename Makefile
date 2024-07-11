all: test
all: vet
all: package
all: package_race
all: run


test: vet
test: base_test

base_test:
	go test ./... -v

vet:
	go vet ./...

package: client

package_race: client_race

client:
	go build -o ./bin/client .

client_race:
	go build --race -o ./bin/client_race .

run:
	./bin/client
