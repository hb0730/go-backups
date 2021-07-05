FROM alpine
WORKDIR /app
RUN mkdir config
ENV FLAG=""
COPY go-backups .
COPY config/application.yml ./config/
ENTRYPOINT /app/go-backups ${FLAG}