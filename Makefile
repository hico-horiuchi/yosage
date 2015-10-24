VERSION     := 0.1.0
GO_BUILDOPT := -ldflags '-s -w -X main.VERSION=$(VERSION)'

gom:
	go get github.com/mattn/gom
	gom install

run:
	gom run *.go ${ARGS}

fmt:
	gom exec goimports -w *.go yosage/*.go

bindata:
	gom exec go-bindata lgtm.png

build: fmt bindata
	gom build $(GO_BUILDOPT) -o bin/yosage *.go

release: fmt
	GOOS=linux GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/yosage$(VERSION).linux-amd64 *.go
	GOOS=darwin GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/yosage$(VERSION).darwin-amd64 *.go
	GOOS=windows GOARCH=amd64 gom build $(GO_BUILDOPT) -o bin/yosage$(VERSION).windows-amd64 *.go

clean:
	rm -f bin/yosage*

link:
	mkdir -p $(GOPATH)/src/github.com/hico-horiuchi
	ln -s $(CURDIR) $(GOPATH)/src/github.com/hico-horiuchi/yosage

unlink:
	rm $(GOPATH)/src/github.com/hico-horiuchi/yosage
	rmdir $(GOPATH)/src/github.com/hico-horiuchi
