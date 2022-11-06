# gin layout

## where is system log ?
- ./logs/gin.log
## where is debug log?
- ./logs/debug.log

# prepare
1. create your env setting
```
cp .env.example .env
```
# start
```
#1.
make build

#2.
gin-layout
```
# cli
- https://github.com/LIOU2021/gin-layout-cli

## ref
- https://github.com/eddycjy/go-gin-example
- https://github.com/haubar/GoMVC

## test
- go test -v
- go test -run Test_CsrfToken -v
- go test -cover -v

## todo list
- create model
- create orm https://ithelp.ithome.com.tw/articles/10245308
- create redis https://github.com/go-redis/redis