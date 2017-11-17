export GOPATH := ${PWD}
APP_DIR := src/webapp

build:
	cd src/webapp && dep ensure
	go build -v webapp

init:
	cd ${APP_DIR} && dep init

reinit:
	cd ${APP_DIR} && rm -r vendor && rm Gopkg.* && dep init

vet:
	go vet ./${APP_DIR}/...
