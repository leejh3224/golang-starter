# Golang starter

golang web project boilerplate with hot reloading, config manager and docker

## Getting Started

1. Install [Docker](https://docs.docker.com/v17.09/engine/installation/)

2. Edit `Makefile`

Please replace `[your-docker-hub-id/your-docker-image-name]` with your docker image name.

```makefile
dev:
    docker run -p 8080:8080 -it --rm --volume=$(PWD):/app [your-docker-hub-id/your-docker-image-name]
```

3. Run command

```bash
make dev
```

4. Good to go, live reloading will work like a charm!
