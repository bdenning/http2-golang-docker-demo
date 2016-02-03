# Golang HTTP/2 and Docker
Use the following to build the demo.
```sh
make binary
make docker-container
docker run --rm -ti bdenning/webservice
```
Then, open Chrome to https://{{container-ip}}:8000/ and use the debug tools (F12) to see that the protocol field shows "h2", indicating HTTP/2.
