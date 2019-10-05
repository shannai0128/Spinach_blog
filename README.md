# Spinach_blog

An example blog of gin framework 


## Installation
```
$ go get github.com/c479096292/Spinach_blog
```
[Download v1.0.0](https://github.com/c479096292/Spinach_blog/releases)


### Required

- Mysql
- Redis

### Ready

Create a **blog database** and import [SQL](https://github.com/c479096292/Spinach_blog/blob/master/blog.sql)

### Conf

You should modify `conf/app.ini`

```
[app]
PageSize = 10
JwtSecret = 233

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
LogMaxSize = 5242880
IpBlackExpire = 60000000000
SecretKey = jeqwijojasofsan8394ty

[server]
# debug or release
SensitiveWords = xx;xx;xx;

...
```

### How To Run
```
$ cd $GOPATH/src/Spinach_blog
then
you should be based on to set environment variable at Spinach_blog/config/db_config.go
```

### Run
```
$ cd $GOPATH/src/Spinach_blog

$ go run main.go 
```


## Features
[ORM_VERSION]
- Swagger
- logging
- Jwt-go
- Gin
- sqlx
- mysql
- Redis

[ORM_VERSION](https://github.com/c479096292/Spinach_blog/tree/feature)