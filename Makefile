GOPATH := $(shell pwd):$(shell pwd)/src/
ENV = GOPATH=${GOPATH}

run:
	cd src/ && ${ENV} go run main.go

build:
	cd src/ && ${ENV} go build main.go

get:
	${ENV} go get github.com/smartystreets/goconvey
	${ENV} go get -u golang.org/x/tools/cmd/goimports

test:
	cd src/great_circle_distance/ && ${ENV} go test -v

webtest:
	cd src/great_circle_distance/ && ${ENV} goconvey
