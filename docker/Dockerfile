FROM scratch
WORKDIR /app/config
WORKDIR /app
ENV FLAG=""
COPY go-backups .
COPY config/application.yml ./config/
ENTRYPOINT /app/go-backups ${FLAG}