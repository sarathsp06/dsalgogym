NODE_ENV=test
GO=go
GOTEST=$(GO) test
NPM=npm

test: npm-install .nyc_output
	$(GOTEST) -race -coverprofile=coverage.txt -covermode=atomic ./...
	$(NPM) test > coverage.lcov

.nyc_output:
	mkdir -p .nyc_output

npm-install:
	$(NPM) install

.PHONY:	test npm-install
