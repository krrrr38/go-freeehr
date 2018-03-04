dep-install:
ifeq ($(shell type dep 2> /dev/null),)
	go get -u github.com/golang/dep/...
endif

dep: dep-install
	dep ensure

lint: dep
	go tool vet -all -printfuncs=Criticalf,Infof,Warningf,Debugf,Tracef .

test: lint
	go test -v freeehr/*.go

clean:
	go clean

.PHONY: dep-install dep lint test clean
