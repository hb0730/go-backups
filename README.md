# go-backups

Compress and upload files to backup server

# command

`go-backups -c="config file"`

## upload config

### git

```yaml
cron:
  blog: "0 0 24 * * *"
blog:
  type: git
  git:
    url:
    username:
    email:
    token:
  uploads:
    # no with extension
    filename: "test"
    dirs:
      - "file path"

```

### oss

#### aliyun

```yaml
cron:
  blog: "0 0 24 * * *"
blog:
  type: aliyun-oss
  aliyun-oss:
    endpoint:
    accessKeyId:
    accessKeySecret:
    bucketName:
  uploads:
    # no with extension
    filename: "test"
    dirs:
      - "file path"
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
      - FLAG="c='/app/config/application.yml'"
    volumes:
      - ./config=/app/config
```