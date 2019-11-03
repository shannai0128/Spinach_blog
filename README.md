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

# bit
LogMaxSize = 5242880
# Nanosecond
IpBlackExpire = 60000000000
SecretKey = 

[server]
# debug or release
SensitiveWords = xx;xx;xx;

...
```

## How to run
```
$ cd $GOPATH/src/Spinach_blog
$ touch run.sh
Copy the following to run.sh
export DBDRIVER=mysql
export DBNAME=spinachBlog
export DBUSER=
export DBPASSWD=

export REDIS_ADDR=127.0.0.1:6379
export REDIS_PASSWORD=
export REDIS_DB=1
go run main.go
```

### Run
```
sh run.sh
```



## Features
[ORM_VERSION]
- gorm
- Swagger
- logging
- Jwt-go
- Gin
- mysql
- Redis

[Basic version](https://github.com/c479096292/Spinach_blog/tree/master)