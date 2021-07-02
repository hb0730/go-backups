FROM alpine
WORKDIR /app
RUN mkdir config
ENV FLAG=""
COPY go-backups .
RUN  ls -al
COPY config config
RUN cd config && ls -al
ENTRYPOINT /app/go-backups ${FLAG}