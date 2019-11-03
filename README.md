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
JwtSecret = 

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

### How To Run
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
$ cd $GOPATH/src/Spinach_blog

$ sh run.sh
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