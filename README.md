# Spinach_blog(ORM version)

An example blog of gin framework 


## Installation
```
$ go get github.com/c479096292/Spinach_blog/tree/feature
```

Or directly [Download v1.1.0](https://github.com/c479096292/Spinach_blog/releases)


### Required

- Mysql
- Redis

### Ready

Create a **blog database**

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

## How to run
```
$ cd $GOPATH/src/Spinach_blog
then
you should be based on to set environment variable at Spinach_blog/config/db_config.go
```

### Run
```
$ go run main.go 
```



## Features

- gorm
- Swagger
- logging
- Jwt-go
- Gin
- mysql
- Redis

[Basic version](https://github.com/c479096292/Spinach_blog/tree/master)