# Spinach_blog

An example blog of gin framework 


## Installation
```
$ go get github.com/c479096292/Spinach_blog
```

## How to run


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

### Run
```
$ cd $GOPATH/src/Spinach_blog

$ go run main.go 
```

Swagger doc

[image](https://i.imgur.com/bVRLTP4.jpg)

## Features

- Swagger
- logging
- Jwt-go
- Gin
- sqlx
- Redis