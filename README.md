Docker Container Number
=======================

If you use docker `--scale` and need to know container number.

Usage
-----

### Docker

Run

    docker run --name dcn --rm -p 127.0.0.1:8080:80 -v /var/run/docker.sock:/var/run/docker.sock -t gupalo/docker-container-number:latest

Call from container

    THREAD_NO=`curl -s http://127.0.0.1:8080/$(hostname)`

### Docker Compose

Add to docker-compose.yaml

    dcn:
        image: 'gupalo/docker-container-number'
        volumes: ['/var/run/docker.sock:/var/run/docker.sock']
        restart: 'always'
        #networks: ['common']

You can add this container to each `docker-compose.yaml` or create global `docker-compose.yaml` with shared network.

Call from container

    export THREAD_NO=`curl -s "http://dcn/$(hostname)"`

Dev
---

Build

    make build

Run

    make run
