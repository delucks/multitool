multitool: *.go
	go test
	export GO111MODULE=on
	go mod tidy
	go build
	strip multitool

install: ~/bin/multitool
	cp multitool ~/bin/
