project = dimm-edac-mon

all: clean mod build-linux

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(project) ./...
mod:
	go get -t -v -d -u ./...
	bash -c 'if [[ ! -f ./go.mod  ]]; then go mod init; fi'
	go mod tidy
	go mod vendor
clean:
	bash -c 'if [[ -f "./$(project)" ]]; then rm -fv $(project); fi'
