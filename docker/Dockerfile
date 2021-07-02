FROM alpine
WORKDIR /app
RUN mkdir config
ENV FLAG=""
COPY go-backups .
RUN  ls -al
COPY config/*.yaml ./config
RUN cd config
RUN ls -al
ENTRYPOINT /app/go-backups ${FLAG}