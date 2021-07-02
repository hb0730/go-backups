# go-backups

Compress and upload files to git server

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
    dirpath:
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