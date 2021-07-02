FROM scratch
WORKDIR /app
RUN mkdir ./config
ENV FLAG=""
COPY ./go-backups .
COPY ./config/*.yaml ./config
RUN  ls
ENTRYPOINT /app/go-backups ${FLAG}