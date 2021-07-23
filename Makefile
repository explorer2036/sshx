all:
	go build -o ssh-term
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ssh-term-linux
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o ssh-term-windows.exe
