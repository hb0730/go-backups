# go-backups

Compress and upload files to backup server

# command

`go-backups -c="config file"`

# application.yml

```yaml
cron:
  # upload name
  blog: "0 0 0 * * *"
# upload name
blog:
  # upload server with type [git|aliyun-oss......]
  type:
  #upload server with type  
  # You can see under the cron  package
  git:
    # file compress type [zip|tar|tar.gz]
    # you  see then compress file
    compress: zip|tar|tar.gz
    url:
    username:
    email:
    token:
  uploads:
    # compress file name with not extension
    filename: test
    dirs:
      #  compression Directory path
      - ""
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