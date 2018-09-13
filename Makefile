multitool: *.go
	export GO111MODULE=on
	go mod tidy
	go build
	strip multitool

install:
	cp multitool ~/bin/
