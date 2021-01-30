go_apps = users/create users/retrieve

users/% : api/users/%.go api/users/dao.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/$@ $< api/users/dao.go

clean:
	rm -rf ./bin

build: $(go_apps)

#deploy: clean build
#	sls deploy --verbose
