build:
	GOOS=linux GOARCH=amd64 go build
	docker build --tag go-file-server .

run:
	docker run -p 8080:80 --rm go-file-server
