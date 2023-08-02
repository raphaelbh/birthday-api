## Development
```docker run --rm --net=host -it -v "$PWD":/usr/src/app -w /usr/src/app golang:1.20 bash``

## Create module
go mod init github.com/raphaelbh/birthday-api

## Install Dependencies
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get -u gorm.io/gorm
go get gorm.io/driver/postgres

# Environment Variables
export NOME_DA_VARIAVEL=valor