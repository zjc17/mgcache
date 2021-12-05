.PHONY: mock test clean lint

mock:
	./script/mock.sh

test: mock
	go test $$(sh -c "go list ./... | grep -v /test") -v -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

lint:
	golint $$(go list ./... | grep -v /test)

clean:
	rm -rf ./test/mock