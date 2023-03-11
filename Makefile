all:
	CGO_ENABLED=1 
	CC=gcc 
	GOOS=windows GOARCH=amd64 
	go build -tags static -ldflags "-s -w"