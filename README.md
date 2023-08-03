## Development
```docker run --rm --net=host -it -v "$PWD":/usr/src/app -w /usr/src/app golang:1.20 bash``

## Create module
go mod init github.com/raphaelbh/birthday-api

## Install Dependencies
go get github.com/gin-gonic/gin
go get github.com/gin-contrib/cors
go get gorm.io/gorm
go get gorm.io/driver/postgres
go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/service/s3

# Environment Variables
export INTERNAL_DSN=""
export ACCESS_KEY=""
export SECRET_KEY=""
