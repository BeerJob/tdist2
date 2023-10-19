# Variables
IMAGE_NAME = continentes

# Comandos
build:
	sudo docker build . -t $(IMAGE_NAME) 

docker-continentes:
	sudo docker run $(IMAGE_NAME)