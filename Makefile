.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/users/create ./api/users
	zip bin/users/create.zip bin/users/create
clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
