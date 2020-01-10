NAME = gupalo/docker-container-number

.PHONY: all build release run

all: build

build:
	docker build -t $(NAME) --rm --pull .

release:
	docker push $(NAME):latest

run:
	docker run --name dcn --rm -p 127.0.0.1:8080:80 -v /var/run/docker.sock:/var/run/docker.sock -t $(NAME):latest
