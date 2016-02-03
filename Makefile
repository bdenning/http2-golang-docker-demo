all: binary docker-container

binary:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' webservice.go

docker-container: webservice
	sudo docker build -t bdenning/webservice .

clean:
	rm -f webservice
	sudo docker rmi bdenning/webservice

.PHONY: binary docker-container clean
