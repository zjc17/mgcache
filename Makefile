.PHONY: mock test clean

mock:
	./script/mock.sh

test: mock
	go test $$(sh -c "go list ./... | grep -v /test") -v -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

clean:
	rm -rf ./test/mock