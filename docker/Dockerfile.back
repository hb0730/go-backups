FROM scratch
WORKDIR /app
ENV FLAG=""
COPY ../go-backups .
COPY ../config .
ENTRYPOINT /app/go-backups ${FLAG}