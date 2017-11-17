export GOPATH := ${PWD}

build:
	cd src/webapp && dep ensure
	go build -v webapp

init:
	cd src/webapp && dep init

reinit:
	cd src/webapp && rm -r vender && rm Gopkg.* && dep init

vet:
	go vet ./src/webapp/...
