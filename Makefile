all: webservice docker

webservice:
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' webservice.go

docker: webservice
	sudo docker build -t bowen/webservice .

clean:
	rm -f webservice
	sudo docker rmi bowen/webservice

.PHONY: webservice docker clean

