dev:
	docker run -p 8080:8080 -it --rm --volume=$(PWD):/app [your-docker-hub-id/your-docker-image-name]

build:
	go build -o bin/main .
