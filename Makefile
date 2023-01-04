install: clean
	rm -rf go.mod go.sum
	go mod init
	go mod tidy

	touch .gitignore

	echo "bin\nbin/todolist" >> .gitignore

	go install github.com/kisielk/errcheck@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go get -u github.com/swaggo/swag/cmd/swag

clean:
	rm -rf go.mod go.sum bin/todolist
	go mod init
	go mod tidy
                       
build:
	go build -o bin/todolist

run: errorcheck test build
	echo "\nStarting an app...\n"
	bin/todolist

test:
	go test ./...

errorcheck:
	go vet ./...
	~/go/bin/staticcheck ./...

format:
	go fmt ./...

swagger-init:
	~/go/bin/swag init
