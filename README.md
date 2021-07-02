# go-backups

Compress and upload files to git server

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
    access_key_id:
    access_key_secret:
    bucket_name:
  uploads:
    # no with extension
    filename: "test"
    dirs:
      - "file path"
```