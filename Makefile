NAME = roustem/hello-docker
VERSION = 1.0

all: image

main:
	env GOOS=linux GOARCH=amd64 go build main.go

image: main
	docker build -t $(NAME):$(VERSION) .

run: image
	docker run --rm -p 8080:8080 $(NAME):$(VERSION)

push: image
	docker push $(NAME):$(VERSION)


