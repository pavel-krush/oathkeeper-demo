all: compile image

compile:
	GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/linux-amd64/server .

image:
	docker build -t authorizer:latest .
