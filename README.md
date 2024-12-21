# bazi-api

## Install Go
#### https://go.dev/dl/

## Download and Install
```
git clone https://github.com/phawitb/bazi-api.git

cd bazi-api
go mod init bazi-api

go get github.com/gin-gonic/gin
go get github.com/tommitoan/bazica
#go mod vendor

# for test
go run bazi_test.go

# run api
go run bazi_api.go

```
```
POST :: http://localhost:8080/api/bazi
BODY :: 
{
    "date": "1993-04-24",
    "time": "06:30",
    "timezone": "Asia/Bangkok",
    "gender": 1
}
```
## Reference
#### https://github.com/tommitoan/bazica
