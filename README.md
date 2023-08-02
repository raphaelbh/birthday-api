## Development
```docker run --rm --net=host -it -v "$PWD":/usr/src/app -w /usr/src/app golang:1.20 bash``


## Create module
go mod init github.com/raphaelbh/birthday-api

## Install Dependencies
go get github.com/gin-gonic/gin