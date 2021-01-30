.PHONY: clean

go_apps = users/create users/retrieve

users/% : src/users/%.go src/users/dao.go
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/$@ $< src/users/dao.go

clean:
	rm -rf ./bin

build: $(go_apps)

#deploy: clean build
#	sls deploy --verbose
