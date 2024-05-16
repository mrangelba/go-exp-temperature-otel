run:
	docker-compose -f docker-compose.yml up

build-docker:
	docker build -f Dockerfile . -t service
