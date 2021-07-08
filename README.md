# go-backups

Compress and upload files to backup server

# command

`go-backups -c="config file"`

# application.yml

```yaml
# cron
cron:
  # upload name and Crone expression
  # see github.com/robfig/cron
  blog: "*/4 * * * * *"
# upload name
blog:
  # type of upload server [git|aliyun-oss|qiniu-oss...]
  type: git
  # type of upload server
  git:
    # compress type [zip|tar...]
    compress: zip
    url:
    username:
    email:
    token:
  # uploads
  uploads:
    # compress name with not extension
    filename:
    # compression directory path
    dirs:
      - ""
```

## compress type

### git

```yaml
compress:
url:
username:
email:
token:
```

### aliyun-oss

```yaml
compress:
endpoint:
accessKeyId:
accessKeySecret:
bucketName:
```

### qiniu-oss

```yaml
compress:
accessKey:
secretKey:
bucket:
regionId:
```

# Docker

```yaml
version: '3.8'
services:
  backups:
    image: hb0730/go-backups
    container_name: backups
    restart: always
    environment:
      - FLAG="-c='/app/config/application.yml'"
    volumes:
      - ./config=/app/config
```

# Support

* git server
* [aliyun-oss](https://www.aliyun.com/product/oss)
* [qiniu-oss](https://www.qiniu.com/products/kodo)