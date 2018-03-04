setup:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif
ifeq ($(shell type golint 2> /dev/null),)
	go get github.com/golang/lint/golint
endif

dep: setup
	dep ensure

lint: dep
	golint --set_exit_status freeehr

vet: dep
	go tool vet -all -printfuncs=Criticalf,Infof,Warningf,Debugf,Tracef .

test: lint vet
	go test -v freeehr/*.go

clean:
	go clean

.PHONY: setup dep lint test clean
