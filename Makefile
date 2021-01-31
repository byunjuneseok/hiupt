.PHONY: clean

lambdas = lambda/userCreate lambda/userRetrieve

lambda/% : src/lambda/%.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$@ $<

clean:
	rm -rf ./bin

build: $(lambdas)

#deploy: clean build
#	sls deploy --verbose
